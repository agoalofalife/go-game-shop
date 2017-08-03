package models

import (
	"fmt"
	"database/sql"
)

type Categories struct {
	id int
	Name string
	Description sql.NullString
	Created_at NullTime
	Update_at NullTime
	Deleted_at NullTime
}


func CategoriesAll() []Categories {
	var (
		err error
		category Categories
		categories []Categories
	)

	results,err := DB.Query("SELECT name, description, created_at, update_at FROM `categories`")
	defer results.Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&category.Name, &category.Description, &category.Created_at, &category.Update_at)
		categories = append(categories, category)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return categories
}