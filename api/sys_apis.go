package api

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"go_grpc_gorm_micro/lib/response"
	"go_grpc_gorm_micro/proto/proto"
	"go_grpc_gorm_micro/service"
	"google.golang.org/protobuf/types/known/anypb"
)

type SysApis struct{}

// 生成curd代码
func (s *SysApis) Create(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.CreateSysApis(req)
	return response.SuccessAny(data), err
}


func (s *SysApis) Delete(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.DeleteSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) DeleteById(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.DeleteByIdSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) Update(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.UpdateSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) Find(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.FindSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) Lists(ctx context.Context, req *proto.Request) (*proto.Responses, error) {
    data, total, err := service.GetListSysApis(req)

	var any = make([]*anypb.Any, len(data))
	for k, r := range data {
		any[k], err = ptypes.MarshalAny(r)
	}

	return response.SuccesssAny(any, total), err
}


