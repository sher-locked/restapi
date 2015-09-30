package main

import (
	"fmt"
	"github.com/Pirate-Labs/restapi/models"
	"net/http"
)

/*Initilise the server */
func main() {
	models.InitDB("sher_locked:kenhar1!1naush@tcp(54.68.25.215:3306)/testdb")

	http.HandleFunc("/users", usersIndex)
	http.ListenAndServe(":3000", nil)
}

func usersIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	users, err := models.AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "%d, %s, %s, %s, %d \n", user.User_Id, user.First_Name, user.Last_Name, user.Wallet, user.Balance)
	}
}
