package dal

import (
	"sfw/biz/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	// DB 数据库
	DB *gorm.DB
	// Executor 查询执行器
	Executor *query.Query
	// DSN 数据库连接字符串
	DSN string
)

// Load 加载数据库
func Load() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	Executor = query.Use(DB)
}
