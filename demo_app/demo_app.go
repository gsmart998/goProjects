package main

import (
	"database/sql"
	"fmt"

	"example.com/packages/db"
	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Dbname)

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
