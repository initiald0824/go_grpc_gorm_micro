package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"go_grpc_gorm_micro/lib/global"
	"time"
)

var err error

// gorm 初始化数据库并产生数据库全局变量
func Gorm() {
	switch global.CURD_CONFIG.System.DbType {
	case "mysql":
		GormMysql()
	default:
		GormMysql()
	}
}

func GormDBTables(db *gorm.DB) {
	// 有错误
	//err := db.AutoMigrate(
	//	sysapi.SysApi{},
	//)
	//if err != nil {
	//	fmt.Println("register table failed")
	//	fmt.Println(err.GetErrors())
	//	os.Exit(0)
	//}
	//fmt.Println("register table success")
}
//// GormMysql 初始化Mysql数据库
func GormMysql() {
	m := global.CURD_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@/" + m.Dbname + "?" + m.Config


	global.CURD_DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		global.CURD_LOG.Error("mysql连接失败", zap.Any("err", err))
		panic(err)
	}
	GormDBTables(global.CURD_DB)
	// See "Important settings" section.
	global.CURD_DB.DB().SetConnMaxLifetime(time.Minute * 3)
	global.CURD_DB.DB().SetMaxOpenConns(m.MaxOpenConns)
	global.CURD_DB.DB().SetMaxIdleConns(m.MaxIdleConns)
	global.CURD_DB.LogMode(m.LogMode)
	global.CURD_LOG.Info("mysql连接成功")
}
