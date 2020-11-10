package service

import (
	"go_grpc_gorm_micro/lib/global"
	"go_grpc_gorm_micro/lib/utils"
	"go_grpc_gorm_micro/model"
	"go_grpc_gorm_micro/proto/proto"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type tplData struct {
	template     *template.Template
	locationPath string
	autoCodePath string
}

func createTemp(tplFileList []string, templateStruct model.TemplateStruct) (err error) {
	basePath := global.CURD_CONFIG.System.Director+"lib/tpl/"
	dataList := make([]tplData, 0, len(tplFileList))
	fileList := make([]string, 0, len(tplFileList))

	//定义一个函数add
	//这个函数要么只有一个返回值，要么有俩返回值且第二个返回值必须是error类型
	add := func (params int)(int, error) {
		return 100+params, nil
	}
	// 定义首字母小写
	case2CamelAndLcfirst := utils.Case2CamelAndLcfirst
	// 定义首字母大写
	case2CamelAndUcfirst := utils.Case2CamelAndUcfirst

	// 根据文件路径生成 tplData 结构体，待填充数据
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value})
	}
	// 生成 *Template, 填充 template 字段
	for index, value := range dataList {

		textByte, err := ioutil.ReadFile(value.locationPath)
		if err != nil {
			return err
		}

		dataList[index].template, err = template.New("").Funcs(template.FuncMap{"add": add,"case2CamelAndLcfirst":case2CamelAndLcfirst,"case2CamelAndUcfirst":case2CamelAndUcfirst}).Parse(string(textByte))
		if err != nil {
			return err
		}
	}

	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.locationPath, basePath)
		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				dataList[index].autoCodePath = global.CURD_CONFIG.System.Director + "/" + trimBase[:lastSeparator] + "/" + templateStruct.TableName + origFileName
			}
		}
	}

	// 生成文件
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}

		if err = value.template.Execute(f, templateStruct); err != nil {
			return err
		}
		_ = f.Close()
	}

	return nil
}

func getAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = getAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}

func getTables(dbName string) (err error, TableNames []model.TableReq) {
	err = global.CURD_DB.Raw("select table_name as table_name from information_schema.tables where TABLE_SCHEMA = ?", dbName).Scan(&TableNames).Error
	return err, TableNames
}

func getColumn(tableName string, dbName string) (err error, Columns []model.ColumnReq) {
	err = global.CURD_DB.Raw("SELECT COLUMN_NAME column_name,DATA_TYPE data_type,COLUMN_KEY column_key,EXTRA extra,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT column_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE TABLE_NAME = ? AND TABLE_SCHEMA = ?", tableName, dbName).Scan(&Columns).Error
	return err, Columns
}

func Generate(tableName string, dbName string) (err error) {
	var templateStruct model.TemplateStruct

	if dbName == "" {
		dbName = global.CURD_CONFIG.Mysql.Dbname
	}

	var tableNames []model.TableReq

	if tableName == "" { // 获取所有表
		err, tableNames = getTables(dbName)
		if err != nil {
			return err
		}
	} else {
		var custableName model.TableReq
		custableName.TableName = tableName
		tableNames = append(tableNames, custableName)
	}

	// 获取 basePath 文件夹下所有tpl文件
	tplFileList, err := getAllTplFile(global.CURD_CONFIG.System.Director+"lib/tpl", nil)
	if err != nil {
		return err
	}

	var req proto.SysApis
	// 遍历所有表，获取表结构
	for _, value := range tableNames{
		// 判断是否有生成过
		req.Path = utils.Lcfirst(utils.Case2Camel(value.TableName))
		req, err := FindByPathSysApi(&req)
		if err != nil && err.Error() != "record not found" {
			global.CURD_LOG.Info(req.Path+" 已经生成过"+err.Error())
			continue
		}

		// 组装数据
		err, columns := getColumn(value.TableName, dbName)
		if err != nil {
			continue
		}
		templateStruct.TableName = value.TableName
		templateStruct.ModelName = utils.Case2Camel(value.TableName)
		templateStruct.RouterName = req.Path
		templateStruct.Fields = columns

		// 生成源码
		err = createTemp(tplFileList, templateStruct)
		if err != nil {
			global.CURD_LOG.Info(req.Path+" 生成源码失败"+err.Error())
			continue
		}
	}

	return nil
}
