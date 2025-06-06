// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/asynccnu/ccnubox-be/be-banner/grpc"
	"github.com/asynccnu/ccnubox-be/be-banner/ioc"
	"github.com/asynccnu/ccnubox-be/be-banner/pkg/grpcx"
	"github.com/asynccnu/ccnubox-be/be-banner/repository/cache"
	"github.com/asynccnu/ccnubox-be/be-banner/repository/dao"
	"github.com/asynccnu/ccnubox-be/be-banner/service"
)

// Injectors from wire.go:

func InitGRPCServer() grpcx.Server {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	bannerDAO := dao.NewMysqlBannerDAO(db)
	cmdable := ioc.InitRedis()
	bannerCache := cache.NewRedisBannerCache(cmdable)
	bannerService := service.NewBannerService(bannerDAO, bannerCache, logger)
	bannerServiceServer := grpc.NewBannerServiceServer(bannerService)
	client := ioc.InitEtcdClient()
	server := ioc.InitGRPCxKratosServer(bannerServiceServer, client, logger)
	return server
}
