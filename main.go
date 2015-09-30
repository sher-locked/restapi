package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	user_id    int16
	first_name string
	last_name  string
	wallet     string
	balance    int16
}

func main() {
	db, err := sql.Open("mysql",
		"sher_locked:kenhar1!1naush@tcp(54.68.25.215:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.user_id, &user.first_name, &user.last_name, &user.wallet, &user.balance)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("%d, %s, %s, %s, %d \n", user.user_id, user.first_name, user.last_name, user.wallet, user.balance)
	}
}
