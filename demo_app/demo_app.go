package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "abul.db.elephantsql.com"
	port     = 5432
	user     = "xpztuuwl"
	password = "xWyIxQxj1AXsK4YfyvIQbUkjxHs1WBt6"
	dbname   = "xpztuuwl"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// add new user
	// insert, err := db.Query("INSERT INTO users (id, first_name, last_name, email) VALUES (3, 'Homer', 'Simpson', 'homer-love-duff@mail.com')")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	raws, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer raws.Close()

	fmt.Println("Successfully connected!")

}
