English | [简体中文](./README.md)简体中文

# go_grpc_gorm_micro
Through go + grpc + proto + Gorm... Quickly generate curd code, and has divided the project structure, see the directory structure for details
The general system architecture is as follows:
<div align=center>
<img src="https://github.com/arrayhua/go_grpc_gorm_micro/curd.png" width=300" height="300" />
</div>

## Overview
Through the data table of the database, quickly generate the code of curd,for example:`mysql`.

## Dependent
- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)
- [cobra](https://github.com/spf13/cobra)
- [protoc-go-inject-tag](https://github.com/favadi/protoc-go-inject-tag)
- [gorm](https://github.com/go-gorm/gorm)

## Directory Structure
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
├── curd.sh                            // todo...
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

## Default sql file
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

## Getting started
Using grpc and protobuf to realize high performance API, The steps are as follows:
1. Design MySQL data structure table
2. Edit go_grpc_gorm_micro/lib/constant/constant.go, Under project file config.yaml, Configure MySQL and configure the connection
3. ./curd.sh tableName 比如`./curd.sh users`