package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// HelloDB струкрута БД tgbotwords.sql
type wordsDB struct {
	id   int
	text string
}

func main() {
	fmt.Println(conn("бот"))
	fmt.Println(conn("Привет"))
	fmt.Println(conn("sds"))
}

func conn(text string) string {
	db, err := sql.Open("mysql", "root:A7bje8971@@/tgbotwords")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch text {
	case "бот":
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
		return randWords(aboutbot)
	case "Пидарас":
		rows, err := db.Query("select * from tgbotwords.dirtywords")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		dirtyWords := []wordsDB{}

		for rows.Next() {
			d := wordsDB{}
			err := rows.Scan(&d.id, &d.text)
			if err != nil {
				fmt.Println(err)
				continue
			}
			dirtyWords = append(dirtyWords, d)
		}
		return randWords(dirtyWords)
	case "Привет":
		rows, err := db.Query("select * from tgbotwords.Hellowords")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		helloWorlds := []wordsDB{}

		for rows.Next() {
			h := wordsDB{}
			err := rows.Scan(&h.id, &h.text)
			if err != nil {
				fmt.Println(err)
				continue
			}
			helloWorlds = append(helloWorlds, h)
		}
		return randWords(helloWorlds)
	default:
		return "Error"
	}
}

func randWords(words []wordsDB) string {
	rand.Seed(time.Now().UnixNano())
	var result = words[rand.Intn(len(words))]
	return result.text
}
