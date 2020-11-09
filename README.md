[English](./README-en.md) | 简体中文

# go_grpc_gorm_micro
通过go+grpc+proto+gorm...快速生成curd代码，并已经划分好了项目结构，详见目录结构

## Overview
通过数据库的数据表快速生成增删改查代码，比如通过：`mysql`.


## Dependent
- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)
- [cobra](https://github.com/spf13/cobra)
- [protoc-go-inject-tag](https://github.com/favadi/protoc-go-inject-tag)
- [gorm](https://github.com/go-gorm/gorm)

## 目录结构
```
├── LICENSE
├── README.md
├── api                                 // API routing
│   ├── sys_api.go
├── cmd                                 // Console commands
│   ├── LICENSE
│   ├── cmd
│   │   ├── curd.go
│   │   ├── root.go
│   │   └── version.go
│   ├── latest_log
│   └── main.go
├── config                             // Structure corresponding to config.yaml configuration file
│   ├── config.go
│   ├── gorm.go
│   ├── system.go
│   └── zap.go
├── config.yaml                        // configuration file
├── curd.sh                            // 未做...
├── go.mod
├── go.sum
├── initialize
│   ├── config.go
│   └── gorm.go
├── latest_log
├── lib
│   ├── middleware                    // middleware, for example:auth、log....
│   ├── model
│   │   └── model.go
│   ├── response
│   │   └── response.go
│   ├── swagger
│   ├── tls
│   │   ├── server.key
│   │   └── server.pem
│   ├── tpl                          // curd template file
│   │   ├── api
│   │   ├── proto
│   │   │   ├── _model.proto.tpl
│   │   │   └── _service.proto.tpl
│   │   └── service
│   └── utils
├── log                             
├── main.go                         // entry file
├── model
│   └── sys_generate.go
├── proto
└── service
    ├── curd.go
    ├── sys_api.go
```

## 默认的sql文件
```
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

## 快速开始
步骤如下：
1. 设计MySQL数据结构表
2. 在项目文件下配置.yaml，配置MySQL、各个目录的更改并配置连接
3. cd cmd && go run main.go curd tableName `比如 go run main.go curd sys_apis `

用grpc和protobuf实现高性能API
1. 例如
protoc --go_out=plugins=grpc:. sys_apis_model.proto 
protoc --go_out=plugins=grpc:. --swagger_out=./proto --grpc-gateway_out=./proto sys_apis_service.proto 
2. main.go注册我们的RPC服务
    pb.RegisterSysApisServiceServer(grpcServer, &api.SysApis{})
3. gRPC转换HTTP，文件地址：/lib/gateway/网关.go
    pb.RegisterSysApisServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)
