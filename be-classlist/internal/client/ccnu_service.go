package client

import (
	"context"
	v1 "github.com/asynccnu/ccnubox-be/be-api/gen/proto/user/v1"
	"github.com/asynccnu/ccnubox-be/be-classlist/internal/conf"
	"github.com/asynccnu/ccnubox-be/be-classlist/internal/errcode"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/pkg/errors"
	"time"
)

type CCNUService struct {
	Cs v1.UserServiceClient
}

func NewCCNUService(cs v1.UserServiceClient) *CCNUService {
	return &CCNUService{Cs: cs}
}

func (c *CCNUService) GetCookie(ctx context.Context, stuID string) (string, error) {
	resp, err := c.Cs.GetCookie(ctx, &v1.GetCookieRequest{
		StudentId: stuID,
	})
	if err != nil {
		return "", errors.Wrap(errcode.ErrCCNULogin, err.Error())
	}
	cookie := resp.Cookie
	return cookie, nil
}

func NewClient(r *etcd.Registry, cf *conf.Registry, logger log.Logger) (v1.UserServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cf.Usersvc), // 需要发现的服务，如果是k8s部署可以直接用服务器本地地址:9001，9001端口是需要调用的服务的端口
		grpc.WithDiscovery(r),
		grpc.WithTimeout(5*time.Second), //由于使用华师的服务,所以设置下超时时间最长为5s
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.NewHelper(logger).WithContext(context.Background()).Errorw("kind", "grpc-client", "reason", "GRPC_CLIENT_INIT_ERROR", "err", err)
		return nil, err
	}
	return v1.NewUserServiceClient(conn), nil
}
