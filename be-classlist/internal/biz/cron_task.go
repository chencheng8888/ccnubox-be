package biz

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/asynccnu/ccnubox-be/common/pkg/cronx"
	"github.com/asynccnu/ccnubox-be/common/pkg/logger"
	"github.com/asynccnu/ccnubox-be/common/tool"
	"golang.org/x/time/rate"
)

type CronTaskExecute struct {
	cronManager *cronx.Manager

	ider      GetStudentIDer
	crawler   ClassCrawler
	ccnu      CCNUServiceProxy
	classRepo ClassRepo
}

func NewCronTaskExecute(log logger.Logger, ider GetStudentIDer, crawler ClassCrawler, ccnu CCNUServiceProxy, classRepo ClassRepo) *CronTaskExecute {
	manager := cronx.NewManager(log)
	cm := &CronTaskExecute{
		cronManager: manager,
		ider:        ider,
		crawler:     crawler,
		ccnu:        ccnu,
		classRepo:   classRepo,
	}
	// 每周日凌晨3点执行一次拉取课表的任务
	// notice: 这里的cron表达式是基于robfig/cron库的，默认不支持秒级别，所以格式是 "分 时 日 月 周"。
	// 如果需要支持秒级别，需要在创建Manager时添加cron.WithSeconds()选项
	// 这里暂时不需要秒级别，所以使用默认的cron表达式格式。
	cm.cronManager.AddTask("pull_class_list", "0 3 * * 0", func(ctx_ context.Context, log_ logger.Logger) {
		pullClassListTask(ctx_, log_, cm.ider, cm.crawler, cm.ccnu, cm.classRepo)
	})
	return cm
}

func (c *CronTaskExecute) Stop() {
	c.cronManager.Stop(context.Background())
}

func pullClassListTask(
	ctx context.Context,
	log logger.Logger,
	ider GetStudentIDer,
	crawler ClassCrawler,
	ccnu CCNUServiceProxy,
	classRepo ClassRepo,
) {
	log.Info("Executing PullClassListTask")

	const (
		pageSize        = 200
		workerCount     = 16
		qps             = 20
		progressLogStep = 200
	)

	// 使用官方令牌桶限流
	limiter := rate.NewLimiter(rate.Limit(qps), qps)

	jobs := make(chan string, 1000)

	var (
		wg        sync.WaitGroup
		processed atomic.Int64
	)

	log.Infof("PullClassListTask started pageSize=%d workerCount=%d qps=%d", pageSize, workerCount, qps)

	// worker pool
	for i := 0; i < workerCount; i++ {
		workerID := i + 1
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			log.Infof("worker-%d started", id)
			defer log.Infof("worker-%d stopped", id)

			for stuID := range jobs {
				if tool.IsGraduated(stuID) {
					// 跳过已经毕业的学生
					log.Infof("worker-%d skipping graduated student %s", id, stuID)
					continue
				}


				// 等待令牌,如果上下文被取消则退出
				if err := limiter.Wait(ctx); err != nil {
					log.Warnf("worker-%d limiter wait canceled: %v", id, err)
					return
				}

				log.Infof("worker-%d processing student %s", id, stuID)

				cookie, err := ccnu.GetCookie(ctx, stuID)
				if err != nil {
					log.Errorf("get cookie failed %s: %v", stuID, err)
					continue
				}
				// 获取当前学年学期
				year, semester := tool.GetCurrentAcademicYearAndSemesterStr(time.Now())

				// 根据学号判断学生类型，获取课表
				var stu Student
				sType := tool.ParseStudentType(stuID)
				switch sType {
				case tool.UnderGraduate:
					stu = &Undergraduate{}
				case tool.PostGraduate:
					stu = &GraduateStudent{}
				default:
					stu = &Undergraduate{}
				}
				infos, scs, _, err := stu.GetClass(ctx, stuID, year, semester, cookie, crawler)
				if err != nil {
					log.Errorf("crawl failed %s: %v", stuID, err)
					continue
				}

				// 保存课表到数据库
				if err := classRepo.SaveClass(ctx, stuID, year, semester, infos, scs); err != nil {
					log.Errorf("save failed %s: %v", stuID, err)
					continue
				}

				count := processed.Add(1)
				if count%progressLogStep == 0 {
					log.Infof("processed %d students so far", count)
				}
			}
		}(workerID)
	}

	// 分页生产任务
	lastStuID := ""
	for {
		stuIDs, err := ider.GetStudentIDs(ctx, lastStuID, pageSize)
		if err != nil {
			log.Errorf("get ids failed: %v", err)
			break
		}
		if len(stuIDs) == 0 {
			log.Info("no more student ids to enqueue")
			break
		}

		firstID := stuIDs[0]
		lastPageID := stuIDs[len(stuIDs)-1]
		log.Infof("enqueueing %d student ids range %s-%s", len(stuIDs), firstID, lastPageID)

		for _, id := range stuIDs {
			jobs <- id
		}

		lastStuID = lastPageID
	}

	close(jobs)
	wg.Wait()

	total := processed.Load()
	log.Infof("Finished PullClassListTask processed=%d lastStuID=%s", total, lastStuID)
}
