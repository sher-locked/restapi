package models

/* User model */
type User struct {
	User_Id    int16
	First_Name string
	Last_Name  string
	Wallet     string
	Balance    int16
}

/* Return all Users from DB */
func AllUsers() ([]*User, error) {

	// retrieve rows from db
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over rows
	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.User_Id, &user.First_Name, &user.Last_Name, &user.Wallet, &user.Balance)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
