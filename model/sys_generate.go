package model


type DBReq struct {
	Database string `json:"database";gorm:"column:database"`
}

type TableReq struct {
	TableName string `json:"tableName"`
}

type ColumnReq struct {
	ColumnName    string `json:"columnName";gorm:"column:column_name"`
	DataType      string `json:"dataType";gorm:"column:data_type"`
	COLUMNKEY     string `json:"columnKey";gorm:"column:column_key"`
	EXTRA      	  string `json:"extra";gorm:"column:extra"`
	DataTypeLong  string `json:"dataTypeLong";gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment";gorm:"column:column_comment"`
}

type TemplateStruct struct {
	ModelName         string  `json:"structName"` // SysApis
	TableName         string  `json:"tableName"`  // sys_apis
	RouterName        string  `json:"routerName"`  // sysApis
	Fields            []ColumnReq `json:"fields"`
}



