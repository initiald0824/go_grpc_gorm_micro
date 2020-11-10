package service

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"go_grpc_gorm_micro/lib/constant"
	"go_grpc_gorm_micro/lib/global"
	"go_grpc_gorm_micro/proto/proto"
)

func Create{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	// validator校验

	//
	if !errors.Is(global.CURD_DB.Where(req).First(&proto.{{.ModelName}}{}).Error, gorm.ErrRecordNotFound) {
		return req, errors.New("重复创建～")
	}

	//
	err := global.CURD_DB.Create(&req).Error
	return req, err
}

func Delete{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.Where(&req).First(&req).Delete(&req).Error
	return req, err
}

func DeleteById{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.First(&req).Delete(&req).Error
	return req, err
}

func Update{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.Update(&req).Error
	return req, err
}

func Find{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.Where(&req).First(&req).Error
	return req, err
}

func GetList{{.ModelName}}(req *proto.Request) (result []*proto.{{.ModelName}}, total int64, err error) {
	if req.PageSize == 0 {
		req.PageSize = constant.PAGESIZE
	}
	if req.Page == 0 {
		req.Page = constant.PAGE
	}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

    unmarshal := &proto.{{.ModelName}}{}
	err = ptypes.UnmarshalAny(req.Query, unmarshal)
	db := global.CURD_DB.Model(&result)
    {{range $i, $v := .Fields}}
    {{if eq $v.DataType "varchar"}}
    if unmarshal.{{case2CamelAndUcfirst $v.ColumnName}} != "" {
        db = db.Where("{{$v.ColumnName}} LIKE ?", "%"+unmarshal.{{case2CamelAndUcfirst $v.ColumnName}}+"%")
    }
    {{else if eq $v.DataType "char"}}
    if unmarshal.{{case2CamelAndUcfirst $v.ColumnName}} != "" {
            db = db.Where("{{$v.ColumnName}} LIKE ?", "%"+unmarshal.{{case2CamelAndUcfirst $v.ColumnName}}+"%")
    }
    {{else if eq $v.DataType "timestamp"}}
    if unmarshal.{{case2CamelAndUcfirst $v.ColumnName}} != nil {
            db = db.Where("{{$v.ColumnName}} = ?", unmarshal.{{case2CamelAndUcfirst $v.ColumnName}})
    }
    {{else if eq $v.DataType "text"}}
    if unmarshal.{{case2CamelAndUcfirst $v.ColumnName}} != "" {
            db = db.Where("{{$v.ColumnName}} LIKE ?", "%"+unmarshal.{{case2CamelAndUcfirst $v.ColumnName}}+"%")
    }{{else}}
    if unmarshal.{{case2CamelAndUcfirst $v.ColumnName}} != 0 {
            db = db.Where("{{$v.ColumnName}} = ?", unmarshal.{{case2CamelAndUcfirst $v.ColumnName}})
    }
    {{end}}
    {{end}}

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


