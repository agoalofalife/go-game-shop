package models

import (
	"fmt"
)

type User struct {
	id         int
	email      string
	name       string
	created_at string
	update_at  string
	deleted_at string
	password   string
	authId     string
}

func UsersAll() User {
	var (
		err  error
		user User
	)

	results, _ := DB.Query("SELECT email, name FROM `users`")

	for results.Next() {
		err = results.Scan(&user.email, &user.name)
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
	id,_ := results.LastInsertId()
	return int(id)
}
