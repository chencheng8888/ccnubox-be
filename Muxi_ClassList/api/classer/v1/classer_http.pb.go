// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.26.1
// source: classer/v1/classer.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationClasserAddClass = "/classer.v1.Classer/AddClass"
const OperationClasserDeleteClass = "/classer.v1.Classer/DeleteClass"
const OperationClasserGetClass = "/classer.v1.Classer/GetClass"
const OperationClasserGetRecycleBinClassInfos = "/classer.v1.Classer/GetRecycleBinClassInfos"
const OperationClasserGetSchoolDay = "/classer.v1.Classer/GetSchoolDay"
const OperationClasserRecoverClass = "/classer.v1.Classer/RecoverClass"
const OperationClasserUpdateClass = "/classer.v1.Classer/UpdateClass"

type ClasserHTTPServer interface {
	// AddClass添加课程
	AddClass(context.Context, *AddClassRequest) (*AddClassResponse, error)
	// DeleteClass删除课程
	DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error)
	// GetClass获取课表
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)
	// GetRecycleBinClassInfos获取回收站的课程(回收站的课程只能保存2个月)
	GetRecycleBinClassInfos(context.Context, *GetRecycleBinClassRequest) (*GetRecycleBinClassResponse, error)
	GetSchoolDay(context.Context, *GetSchoolDayReq) (*GetSchoolDayResp, error)
	// RecoverClass恢复课程
	RecoverClass(context.Context, *RecoverClassRequest) (*RecoverClassResponse, error)
	// UpdateClass更新课程
	UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error)
}

func RegisterClasserHTTPServer(s *http.Server, srv ClasserHTTPServer) {
	r := s.Route("/")
	r.GET("/class/get/{stu_id}/{year}/{semester}", _Classer_GetClass0_HTTP_Handler(srv))
	r.POST("/class/add", _Classer_AddClass0_HTTP_Handler(srv))
	r.DELETE("/class/delete/{stuId}/{year}/{semester}/{id}", _Classer_DeleteClass0_HTTP_Handler(srv))
	r.PUT("/class/update", _Classer_UpdateClass0_HTTP_Handler(srv))
	r.GET("/class/recycle/{stuId}/{year}/{semester}", _Classer_GetRecycleBinClassInfos0_HTTP_Handler(srv))
	r.PUT("/class/recover", _Classer_RecoverClass0_HTTP_Handler(srv))
	r.GET("/class/day/get", _Classer_GetSchoolDay0_HTTP_Handler(srv))
}

func _Classer_GetClass0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetClassRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserGetClass)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetClass(ctx, req.(*GetClassRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetClassResponse)
		return ctx.Result(200, reply)
	}
}

func _Classer_AddClass0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddClassRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserAddClass)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddClass(ctx, req.(*AddClassRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddClassResponse)
		return ctx.Result(200, reply)
	}
}

func _Classer_DeleteClass0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteClassRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserDeleteClass)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteClass(ctx, req.(*DeleteClassRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteClassResponse)
		return ctx.Result(200, reply)
	}
}

func _Classer_UpdateClass0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateClassRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserUpdateClass)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateClass(ctx, req.(*UpdateClassRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateClassResponse)
		return ctx.Result(200, reply)
	}
}

func _Classer_GetRecycleBinClassInfos0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRecycleBinClassRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserGetRecycleBinClassInfos)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRecycleBinClassInfos(ctx, req.(*GetRecycleBinClassRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRecycleBinClassResponse)
		return ctx.Result(200, reply)
	}
}

func _Classer_RecoverClass0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RecoverClassRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserRecoverClass)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RecoverClass(ctx, req.(*RecoverClassRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecoverClassResponse)
		return ctx.Result(200, reply)
	}
}

func _Classer_GetSchoolDay0_HTTP_Handler(srv ClasserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSchoolDayReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClasserGetSchoolDay)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSchoolDay(ctx, req.(*GetSchoolDayReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSchoolDayResp)
		return ctx.Result(200, reply)
	}
}

type ClasserHTTPClient interface {
	AddClass(ctx context.Context, req *AddClassRequest, opts ...http.CallOption) (rsp *AddClassResponse, err error)
	DeleteClass(ctx context.Context, req *DeleteClassRequest, opts ...http.CallOption) (rsp *DeleteClassResponse, err error)
	GetClass(ctx context.Context, req *GetClassRequest, opts ...http.CallOption) (rsp *GetClassResponse, err error)
	GetRecycleBinClassInfos(ctx context.Context, req *GetRecycleBinClassRequest, opts ...http.CallOption) (rsp *GetRecycleBinClassResponse, err error)
	GetSchoolDay(ctx context.Context, req *GetSchoolDayReq, opts ...http.CallOption) (rsp *GetSchoolDayResp, err error)
	RecoverClass(ctx context.Context, req *RecoverClassRequest, opts ...http.CallOption) (rsp *RecoverClassResponse, err error)
	UpdateClass(ctx context.Context, req *UpdateClassRequest, opts ...http.CallOption) (rsp *UpdateClassResponse, err error)
}

type ClasserHTTPClientImpl struct {
	cc *http.Client
}

func NewClasserHTTPClient(client *http.Client) ClasserHTTPClient {
	return &ClasserHTTPClientImpl{client}
}

func (c *ClasserHTTPClientImpl) AddClass(ctx context.Context, in *AddClassRequest, opts ...http.CallOption) (*AddClassResponse, error) {
	var out AddClassResponse
	pattern := "/class/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationClasserAddClass))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClasserHTTPClientImpl) DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...http.CallOption) (*DeleteClassResponse, error) {
	var out DeleteClassResponse
	pattern := "/class/delete/{stuId}/{year}/{semester}/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClasserDeleteClass))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClasserHTTPClientImpl) GetClass(ctx context.Context, in *GetClassRequest, opts ...http.CallOption) (*GetClassResponse, error) {
	var out GetClassResponse
	pattern := "/class/get/{stu_id}/{year}/{semester}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClasserGetClass))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClasserHTTPClientImpl) GetRecycleBinClassInfos(ctx context.Context, in *GetRecycleBinClassRequest, opts ...http.CallOption) (*GetRecycleBinClassResponse, error) {
	var out GetRecycleBinClassResponse
	pattern := "/class/recycle/{stuId}/{year}/{semester}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClasserGetRecycleBinClassInfos))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClasserHTTPClientImpl) GetSchoolDay(ctx context.Context, in *GetSchoolDayReq, opts ...http.CallOption) (*GetSchoolDayResp, error) {
	var out GetSchoolDayResp
	pattern := "/class/day/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClasserGetSchoolDay))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClasserHTTPClientImpl) RecoverClass(ctx context.Context, in *RecoverClassRequest, opts ...http.CallOption) (*RecoverClassResponse, error) {
	var out RecoverClassResponse
	pattern := "/class/recover"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationClasserRecoverClass))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClasserHTTPClientImpl) UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...http.CallOption) (*UpdateClassResponse, error) {
	var out UpdateClassResponse
	pattern := "/class/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationClasserUpdateClass))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
