package response

import (
	protobuf "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go_grpc_gorm_micro/proto/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	ERROR   = 400
	SUCCESS = 200
)

type Response struct {
	Code int32          `json:"code"`
	Message  string      `json:"message"`
	Data interface{} 	`json:"data"`
	Meta interface{} 	`json:"meta"`
}

func (err *Response) Error() string {
	return err.Message
}

func FailWithMessage(message string) *Response {
	return &Response{Code: ERROR, Message:message, Data: map[string]interface{}{}, Meta:map[string]interface{}{}}
}

func Success() *Response {
	return &Response{Code: SUCCESS, Message:"请求成功", Data: map[string]interface{}{}, Meta:map[string]interface{}{}}
}

func FailAny(data protobuf.Message) *proto.Response {
	any, _ := ptypes.MarshalAny(data)

	meta := &proto.Meta{Total:1} // 需要进行初始化,初始化为0显示不出来？？？

	rsp := &proto.Response{
		Code: ERROR,
		Message: "请求出错",
		Data: any,
		Meta: meta,
	}
	return rsp
}

func SuccessAny(data protobuf.Message) *proto.Response {
	any, _ := ptypes.MarshalAny(data)

	meta := &proto.Meta{Total:1} // 需要进行初始化,初始化为0显示不出来？？？

	rsp := &proto.Response{
		Code: SUCCESS,
		Message: "请求成功",
		Data: any,
		Meta: meta,
	}
	return rsp
}

func SuccesssAny(any []*anypb.Any, total int64) *proto.Responses {
	meta := &proto.Meta{Total:total} // 需要进行初始化,初始化为0显示不出来？？？

	rsps := &proto.Responses{
		Code: SUCCESS,
		Message: "请求成功",
		Data: any,
		Meta: meta,
	}
	return rsps
}






