package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-code-generate/sqlc-demo/tutorial"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("mysql", "root:a1131657944@/sqlc-test?parseTime=true")
	if err != nil {
		fmt.Println("db connect error")
		return
	}

	model := tutorial.New(db)
	tmp1 := tutorial.CreateAuthorParams{
		Name:    "shubin",
		Country: "china",
		Age:     24,
	}

	tmp2 := tutorial.CreateAuthorParams{
		Name:    "zhao",
		Country: "indian",
		Age:     19,
	}

	result, err := model.CreateAuthor(ctx, tmp1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.RowsAffected())

	result, err = model.CreateAuthor(ctx, tmp2)
	if err != nil {
		fmt.Println(err)
		return
	}

	author, err := model.GetAuthor(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(author)

	authorList, err := model.ListAuthors(ctx)
	if err != nil {
		return
	}
	fmt.Println(authorList)

}
