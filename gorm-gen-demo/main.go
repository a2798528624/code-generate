package main

// gorm gen configure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gen"
)

const MySQLDSN = "root:a1131657944@tcp(127.0.0.1:3306)/sqlc-test?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		/* ModelPkgPath: "dal/model"*/
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(connectDB(MySQLDSN))

	g.ApplyBasic(g.GenerateAllTable()...)

	// 执行并生成代码
	g.Execute()
}
