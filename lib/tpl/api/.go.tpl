package api

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"go_grpc_gorm_micro/lib/response"
	"go_grpc_gorm_micro/proto/proto"
	"go_grpc_gorm_micro/service"
	"google.golang.org/protobuf/types/known/anypb"
)

type {{.ModelName}} struct{}

// 生成curd代码
func (s *{{.ModelName}}) Create(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Create{{.ModelName}}(req)
	return response.SuccessAny(data), err
}


func (s *{{.ModelName}}) Delete(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Delete{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) DeleteById(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.DeleteById{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) Update(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Update{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) Find(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Find{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) Lists(ctx context.Context, req *proto.Request) (*proto.Responses, error) {
    data, total, err := service.GetList{{.ModelName}}(req)

	var any = make([]*anypb.Any, len(data))
	for k, r := range data {
		any[k], err = ptypes.MarshalAny(r)
	}

	return response.SuccesssAny(any, total), err
}


