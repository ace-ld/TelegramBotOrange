package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// HelloDB струкрута БД tgbotwords.sql
type wordsDB struct {
	id   int
	text string
}

func main() {
	db, err := sql.Open("mysql", "root:A7bje8971@@/tgbotwords")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from tgbotwords.aboutbot")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	aboutbot := []wordsDB{}

	for rows.Next() {
		a := wordsDB{}
		err := rows.Scan(&a.id, &a.text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		aboutbot = append(aboutbot, a)
	}
	for _, value := range aboutbot {
		fmt.Println(value.text)
	}
}
