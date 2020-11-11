package main

import (
	"bufio"
	"github.com/gookit/color"
	"io"
	"os"
	"strings"
	"unicode"
	"fmt"
)

// 因为目前找不到好办法替换proto中的多余struct引起的报错，所以这只是暂时方法，欢迎大家有好思路提醒一下
func main() {
	tableName := os.Args[1]
	if tableName == "" {
		color.Info.Println("err 参数数据表必传")
		return
	}
	err := fileChange(tableName)
	if err != nil {
		color.Info.Println(err)
		return
	}

	fmt.Println("FileChange success")
}

// 文件中字符串替换
func fileChange(tableName string) (err error) {
	caseTableName := case2CamelAndUcfirst(tableName)
	path, _ := os.Getwd()
	path = path+"/../"

	// 替换main.go
	err = changeFileChar(path+"main.go",
		"pb.RegisterSysApisServiceServer(grpcServer, &api.SysApis{})",
		"pb.RegisterSysApisServiceServer(grpcServer, &api.SysApis{})\n	pb.Register"+caseTableName+"ServiceServer(grpcServer, &api."+caseTableName+"{})",
	)

	// 替换lib/gateway/gateway.go
	err = changeFileChar(path+"lib/gateway/gateway.go",
		"err = pb.RegisterSysApisServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)",
		"err = pb.RegisterSysApisServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)\n	err = pb.Register"+caseTableName+"ServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)",
	)

	// 替换proto
	err = changeFileChar(path+"proto/proto/"+tableName+"_model.pb.go",
		"type Timestamp = timestamppb.Timestamp",
		"",
	)
	err = changeFileChar(path+"proto/proto/"+tableName+"_service.pb.go",
		"_ \"api\"",
		"",
	)

	return err
}


func changeFileChar(fileName string, oldString string, newString string) (err error) {
	in, err := os.Open(fileName)
	if err != nil {
		os.Exit(-1)
	}
	defer in.Close()

	outFileName := fileName+".bak"
	out, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(in)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), oldString, newString, -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			os.Exit(-1)
		}
		index++
	}
	os.Remove(fileName)
	os.Rename(outFileName, fileName)
	return err
}


// 下划线写法转为驼峰写法并且首字母大写
func case2CamelAndUcfirst(name string) string {
	return ucfirst(case2Camel(name))
}

// 首字母大写
func ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 下划线写法转为驼峰写法
func case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
