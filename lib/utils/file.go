package utils

import (
	"bufio"
	"go.uber.org/zap"
	"go_grpc_gorm_micro/lib/global"
	"io"
	"os"
	"strings"
)

// 文件中字符的替换
func ChangeFileChar(fileName string, oldString string, newString string) (err error) {
	in, err := os.Open(fileName)
	if err != nil {
		global.CURD_LOG.Error("open file fail:"+ fileName, zap.Any(" error:", err))
		os.Exit(-1)
	}
	defer in.Close()

	outFileName := fileName+".bak"
	out, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		global.CURD_LOG.Error("Open write file fail:"+ fileName, zap.Any(" error:", err))
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
			global.CURD_LOG.Error("read err:"+ fileName, zap.Any(" error:", err))
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), oldString, newString, -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			global.CURD_LOG.Error("write to file fail:"+ fileName, zap.Any(" error:", err))
			os.Exit(-1)
		}
		index++
	}
	os.Remove(fileName)
	os.Rename(outFileName, fileName)
	return err
}
