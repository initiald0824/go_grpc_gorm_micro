package service

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"go_grpc_gorm_micro/lib/constant"
	"go_grpc_gorm_micro/lib/global"
	"go_grpc_gorm_micro/proto/proto"
)

func CreateSysApis(req *proto.SysApis) (*proto.SysApis, error) {
	// validator校验
	data := req

	//
	if !errors.Is(global.CURD_DB.Where(&req).First(&proto.SysApis{}).Error, gorm.ErrRecordNotFound) {
	//if !errors.Is(global.CURD_DB.Where("path = ? AND method = ?", req.Path, req.Method).First(&proto.SysApis{}).Error, gorm.ErrRecordNotFound) {
		return data, errors.New("重复创建～")
	}

	//
	err := global.CURD_DB.Create(&req).Error
	return data, err
}

func DeleteSysApis(req *proto.SysApis) (*proto.SysApis, error) {
	err := global.CURD_DB.Where(&req).First(&req).Delete(&req).Error
	return req, err
}

func DeleteByIdSysApis(req *proto.SysApis) (*proto.SysApis, error) {
	err := global.CURD_DB.First(&req).Delete(&req).Error
	return req, err
}

func UpdateSysApis(req *proto.SysApis) (*proto.SysApis, error) {
	err := global.CURD_DB.Update(&req).Error
	return req, err
}

func FindSysApis(req *proto.SysApis) (*proto.SysApis, error) {
	err := global.CURD_DB.Where(&req).First(&req).Error
	return req, err
}

func GetListSysApis(req *proto.Request) (result []*proto.SysApis, total int64, err error) {
	if req.PageSize == 0 {
		req.PageSize = constant.PAGESIZE
	}
	if req.Page == 0 {
		req.Page = constant.PAGE
	}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

    unmarshal := &proto.SysApis{}
	err = ptypes.UnmarshalAny(req.Query, unmarshal)
	db := global.CURD_DB.Model(&result)

	if unmarshal.Path != "" {
		db = db.Where("path LIKE ?", "%"+unmarshal.Path+"%")
	}

	if unmarshal.Description != "" {
		db = db.Where("description LIKE ?", "%"+unmarshal.Description+"%")
	}

	if unmarshal.Method != "" {
		db = db.Where("method = ?", unmarshal.Method)
	}

	if unmarshal.ApiGroup != "" {
		db = db.Where("api_group = ?", unmarshal.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return result, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if req.OrderKey != "" {
			var OrderStr string
			if req.OrderDesc != "" {
				OrderStr = req.OrderKey + " desc"
			} else {
				OrderStr = req.OrderKey
			}
			err = db.Order(OrderStr).Find(&result).Error
		} else {
			err = db.Order("id").Find(&result).Error
		}
	}

	return result, total, err
}

func FindByPathSysApi(req *proto.SysApis) (*proto.SysApis, error) {
	db := global.CURD_DB.Model(&req)

	if req.Path != "" {
		db = db.Where("path LIKE ?", req.Path+"%")
	}
	err := db.First(&req).Error
	return req, err
}

