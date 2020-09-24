package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:A7bje8971@@/productdb") // подключение к БД

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into productdb.Products (model, company, price) values (?, ?, ?)", "iPhone X", "Apple", 72000) // добавление информации в столбцы
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
