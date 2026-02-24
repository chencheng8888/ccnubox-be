package grpc

import (
	"context"

	"github.com/asynccnu/ccnubox-be/be-content/domain"
	contentv1 "github.com/asynccnu/ccnubox-be/common/api/gen/proto/content/v1"
)

func (c *ContentServiceServer) GetSemester(ctx context.Context, in *contentv1.GetSemesterRequest) (*contentv1.GetSemesterResponse, error) {
	semester, err := c.svcSemester.Get(ctx, in.GetDate())
	if err != nil {
		return nil, err
	}
	return &contentv1.GetSemesterResponse{
		Semester: semester,
	}, nil
}

func (c *ContentServiceServer) SaveSemester(ctx context.Context, in *contentv1.SaveSemesterRequest) (*contentv1.SaveSemesterResponse, error) {
	semester := domain.Semester{
		Semester:  in.GetSemester().GetSemester(),
		StartDate: in.Semester.GetStartDate(),
		EndDate:   in.Semester.GetEndDate(),
	}
	err := c.svcSemester.Save(ctx, &semester)
	if err != nil {
		return nil, err
	}
	return &contentv1.SaveSemesterResponse{}, nil

}
