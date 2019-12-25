package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
)

func main() {
	db,err := sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/tree_hole?charset=utf8")
	if err != nil {
		panic(err)
	}

	smt,err := db.Prepare("select * from t_user")
	if err != nil {
		panic(err)
	}

	res,err := smt.Query()
	if err != nil {
		panic(err)
	}

	cols,err := res.ColumnTypes()
	if err != nil {
		panic(err)
	}

	for i := 0;i< len(cols);i++  {
		fmt.Println(cols[i].Name())
	}

	for i:=0;i<100 ;i++  {
		num := rand.Intn(10)
		fmt.Println(num)
	}
}