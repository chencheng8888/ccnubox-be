package content

import (
	"time"

	"github.com/asynccnu/ccnubox-be/bff/errs"
	"github.com/asynccnu/ccnubox-be/bff/pkg/ginx"
	"github.com/asynccnu/ccnubox-be/bff/web"
	contentv1 "github.com/asynccnu/ccnubox-be/common/api/gen/proto/content/v1"
	"github.com/gin-gonic/gin"
)

func (h *ContentHandler) RegisterSemesterRoute(group *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	sg := group.Group("/semester")
	sg.GET("/getSemester", authMiddleware, ginx.Wrap(h.GetSemester))
	sg.POST("/saveSemester", authMiddleware, ginx.WrapReq(h.SaveSemester))
	sg.GET("/getSemesterList", authMiddleware, ginx.Wrap(h.GetSemesterList))
}

// GetSemester 获取当前所属学期
// @Summary 获取当前所属学期
// @Description 获取当前所属学期
// @Param Authorization header string true "Bearer Token"
// @Tags semester
// @Success 200 {object} web.Response{data=GetSemesterResponse} "成功"
// @Router /semester/getSemester [get]
func (h *ContentHandler) GetSemester(ctx *gin.Context) (web.Response, error) {
	r := &contentv1.GetSemesterRequest{Date: time.Now().Format("2006-01-02")}
	resp, err := h.contentClient.GetSemester(ctx, r)
	if err != nil {
		return web.Response{}, errs.GET_SEMESTER_ERROR(err)
	}
	data := GetSemesterResponse{Semester: resp.Semester}
	return web.Response{
		Msg:  "Success",
		Data: data,
	}, nil
}

// SaveSemester 保存学期信息
// @Summary 保存学期信息
// @Description 保存学期信息
// @Param Authorization header string true "Bearer Token"
// @Param request body SaveSemesterRequest true "保存学期信息请求参数"
// @Tags semester
// @Success 200 {object} web.Response{} "成功"
// @Router /semester/saveSemester [post]
func (h *ContentHandler) SaveSemester(ctx *gin.Context, req SaveSemesterRequest) (web.Response, error) {
	r := &contentv1.SaveSemesterRequest{
		Semester: &contentv1.Semester{Semester: req.Semester, StartDate: req.StartDate, EndDate: req.EndDate},
	}

	_, err := h.contentClient.SaveSemester(ctx, r)
	if err != nil {
		return web.Response{}, errs.SAVE_SEMESTER_ERROR(err)
	}
	return web.Response{
		Msg: "Success",
	}, nil
}

// GetSemesterList 获取所有学期信息
// @Summary 获取所有学期信息
// @Description 获取所有学期信息
// @Param Authorization header string true "Bearer Token"
// @Tags semester
// @Success 200 {object} web.Response{data=GetSemesterListResponse} "成功"
// @Router /semester/getSemesterList [get]
func (h *ContentHandler) GetSemesterList(ctx *gin.Context) (web.Response, error) {
	resp, err := h.contentClient.GetSemesterList(ctx, &contentv1.GetSemesterListRequest{})
	if err != nil {
		return web.Response{}, err
	}

	semesters := make([]Semester, 0, len(resp.Semesters))
	for _, s := range resp.Semesters {
		semesters = append(semesters, Semester{
			Semester:  s.GetSemester(),
			StartDate: s.GetStartDate(),
			EndDate:   s.GetEndDate(),
		})
	}

	return web.Response{
		Msg:  "Success",
		Data: semesters,
	}, nil
}
