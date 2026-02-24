package service

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/asynccnu/ccnubox-be/be-content/domain"
	"github.com/asynccnu/ccnubox-be/be-content/repository"
	"github.com/asynccnu/ccnubox-be/be-content/repository/model"
	contentv1 "github.com/asynccnu/ccnubox-be/common/api/gen/proto/content/v1"
	"github.com/asynccnu/ccnubox-be/common/pkg/errorx"
	"github.com/asynccnu/ccnubox-be/common/pkg/logger"
)

var (
	GET_SEMESTER_ERROR  = errorx.FormatErrorFunc(contentv1.ErrorGetSemesterError("获取学期列表失败"))
	SAVE_SEMESTER_ERROR = errorx.FormatErrorFunc(contentv1.ErrorSaveSemesterError("保存学期信息失败"))
)

type SemesterService interface {
	Get(ctx context.Context, t string) (string, error)
	Save(ctx context.Context, s *domain.Semester) error
}

type semesterService struct {
	repo repository.ContentRepo[model.Semester]
	l    logger.Logger
}

func NewSemesterService(repo repository.ContentRepo[model.Semester], l logger.Logger) SemesterService {
	return &semesterService{repo: repo, l: l}
}

func (se *semesterService) Get(ctx context.Context, date string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", GET_SEMESTER_ERROR(err)
	}
	//获取所有学期列表
	semesters, err := se.repo.GetList(ctx)
	if err != nil {
		return "", GET_SEMESTER_ERROR(err)
	}

	//命中学期的情况
	for _, semester := range semesters {
		if !t.Before(semester.StartDate) && !t.After(semester.EndDate) {
			return semester.Semester, nil
		}
	}

	//未命中某学期，取距离区间最近的学期
	var closestSemester string
	minDistance := time.Duration(math.MaxInt64)
	var distance time.Duration
	for _, semester := range semesters {
		switch {
		case t.Before(semester.StartDate):
			distance = semester.StartDate.Sub(t)
		case t.After(semester.EndDate):
			distance = t.Sub(semester.EndDate)
		default:
			distance = 0
		}
		if distance < minDistance {
			minDistance = distance
			closestSemester = semester.Semester
		}
	}
	return closestSemester, nil
}

func (se *semesterService) Save(ctx context.Context, s *domain.Semester) error {
	//把日期字符串转换成时间类型
	startDate, err := time.Parse("2006-01-02", s.StartDate)
	if err != nil {
		return SAVE_SEMESTER_ERROR(err)
	}
	endDate, err := time.Parse("2006-01-02", s.EndDate)
	if err != nil {
		return SAVE_SEMESTER_ERROR(err)
	}

	record, err := se.repo.Get(ctx, "semester", s.Semester)
	//记录不存在：创建新记录
	if errors.Is(err, repository.ErrRecordNotFound) {
		modelSemester := &model.Semester{
			Semester:  s.Semester,
			StartDate: startDate,
			EndDate:   endDate,
		}
		err = se.repo.Save(ctx, modelSemester)
		if err != nil {
			return SAVE_SEMESTER_ERROR(err)
		}
		return nil
	}
	if err != nil {
		return SAVE_SEMESTER_ERROR(err)
	}

	//如果存在记录，更新记录
	record.Semester = s.Semester
	record.StartDate = startDate
	record.EndDate = endDate
	record.UpdatedAt = time.Now()

	err = se.repo.Save(ctx, record)
	if err != nil {
		return SAVE_SEMESTER_ERROR(err)
	}
	return nil
}
