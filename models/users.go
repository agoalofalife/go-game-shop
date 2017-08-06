package models

import (
	"fmt"
)

type User struct {
	Id         int
	Email      string
	Name       string
	Created_at string
	Update_at  string
	Deleted_at string
	Password   string
	authId     string
}

func UsersAll() User {
	var (
		err  error
		user User
	)

	results, _ := DB.Query("SELECT email, name FROM `users`")

	for results.Next() {
		err = results.Scan(&user.Email, &user.Name)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return user
}

func UserCreate(parameters map[string]interface{}) int {

	stmt, _ := DB.Prepare("INSERT users SET email=?, name=?,password=?")
	results, _ := stmt.Exec(parameters["email"], parameters["name"], parameters["password"])
	results.RowsAffected()
	id, _ := results.LastInsertId()
	return int(id)
}

func UserGetByEmail(email string) User {
	var (
		user User
	)

	results := DB.QueryRow("SELECT id, password FROM `users` WHERE email = ?", email)

	err := results.Scan(&user.Id, &user.Password)
	if err != nil {
		fmt.Print(err.Error())
	}
	return user
}
