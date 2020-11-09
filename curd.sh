#!/bin/bash
read tableName
cd cmd/
go run main.go curd -t ${tableName}