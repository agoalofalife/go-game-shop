package models

import "fmt"

type News struct {
	Id          int
	Title       string
	Image       Attachments
	Description string
	Created_at  NullTime
	Update_at   NullTime
	Deleted_at  NullTime
}

func NewsAll() []News {
	var (
		err        error
		new        News
		news       []News
		attachment Attachments
	)

	results, err := DB.Query("SELECT news.id, news.title, news.description, news.created_at, attachments.path FROM `news` LEFT JOIN attachments ON news.image = attachments.id")

	defer results.Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&new.Id, &new.Title, &new.Description, &new.Created_at, &attachment.Path)
		new.Image = attachment
		news = append(news, new)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return news
}

func NewById(id int) News {
	var (
		err        error
		new        News
		attachment Attachments
	)

	results := DB.QueryRow("SELECT news.id, news.title, news.description, attachments.path FROM `news` LEFT JOIN attachments ON news.image = attachments.id WHERE  news.id = ? ", id)
	if err != nil {
		panic(err.Error())
	}
	err = results.Scan(&new.Id, &new.Title, &new.Description, &attachment.Path)
	new.Image = attachment
	if err != nil {
		fmt.Print(err.Error())
	}

	return new
}
