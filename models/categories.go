package models

import (
	"fmt"
	"database/sql"
)

type Categories struct {
	Id int
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

	results,err := DB.Query("SELECT id,name, description, created_at, update_at FROM `categories`")
	defer results.Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&category.Id, &category.Name, &category.Description, &category.Created_at, &category.Update_at)
		categories = append(categories, category)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return categories
}

func CategoryName(id int) Categories{
	var (
		err error
		category Categories
	)

	results := DB.QueryRow("SELECT name FROM `categories` WHERE id = ?", id)

	if err != nil {
		panic(err.Error())
	}
	err = results.Scan(&category.Name)
	if err != nil {
		fmt.Print(err.Error())
	}

	return category
}