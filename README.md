[English](./README-en.md) | 简体中文

## 快速开始
用grpc和protobuf实现高性能API，步骤如下：
1. git clone https://github.com/arrayhua/go_grpc_gorm_micro.git && 设计MySQL数据结构表
2. 修改go_grpc_gorm_micro/lib/constant/constant.go,在项目文件下配置config.yaml，配置MySQL、各个目录的更改并配置连接
3. ./curd.sh tableName 比如`./curd.sh users`

# go_grpc_gorm_micro
通过go+grpc+proto+gorm...快速生成curd代码，并已经划分好了项目结构，详见目录结构
大致系统架构图如下
<div align=center>
<img src="https://img-blog.csdnimg.cn/20201111172903218.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTEzMzcyODA=,size_16,color_FFFFFF,t_70"/>
</div>

## Overview
通过数据库的数据表快速生成增删改查代码，比如通过：`mysql`.


## Dependent
```
ps:不要认为依赖项很多，觉得是需要全部掌握才可以上手哦。其实只需要懂MVC和GO语言基础即可完成业务需求。
```
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```
