package main

import (
	"context"
	"fmt"
	"go-code-generate/gorm-gen-demo/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	ctx := context.Background()
	db := connectDB(MySQLDSN)
	query.SetDefault(db)

	result, err := query.Author.WithContext(ctx).Where(query.Author.ID.Eq(1)).Find()
	if err != nil {
		panic(err)
	}
	for _, item := range result {
		fmt.Println(item)
	}
}
