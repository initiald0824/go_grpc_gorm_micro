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
	Msg  string      	`json:"msg"`
	Data interface{} 	`json:"data"`
	Meta interface{} 	`json:"meta"`
}

func (err *Response) Error() string {
	return err.Msg
}

func FailWithMessage(msg string) *Response {
	return &Response{Code: ERROR, Msg:msg, Data: map[string]interface{}{}, Meta:map[string]interface{}{}}
}

func Success() *Response {
	return &Response{Code: SUCCESS, Msg:"请求成功", Data: map[string]interface{}{}, Meta:map[string]interface{}{}}
}

func SuccessAny(data protobuf.Message) *proto.Response {
	any, _ := ptypes.MarshalAny(data)

	meta := &proto.Meta{Total:1} // 需要进行初始化,初始化为0显示不出来？？？

	rsp := &proto.Response{
		Code:  SUCCESS,
		Msg: "请求成功",
		Data: any,
		Meta: meta,
	}
	return rsp
}

func SuccesssAny(data protobuf.Message) *proto.Responses {
	//any, _ := ptypes.MarshalAny(data)
	var any = []*anypb.Any{}

	meta := &proto.Meta{Total:1} // 需要进行初始化,初始化为0显示不出来？？？

	rsps := &proto.Responses{
		Code:  SUCCESS,
		Msg: "请求成功",
		Data: any,
		Meta: meta,
	}
	return rsps
}






