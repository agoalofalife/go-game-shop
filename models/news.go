package models

import "fmt"

type News struct {
	Id int
	Title string
	Image Attachments
	Description string
	Created_at NullTime
	Update_at NullTime
	Deleted_at NullTime
}

func NewsAll() []News {
	var (
		err error
		new News
		news []News
		attachment Attachments
	)

	results,err := DB.Query("SELECT news.id, news.title, news.description, attachments.path FROM `news` LEFT JOIN attachments ON news.image = attachments.id")

	defer results.Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&new.Id, &new.Title, &new.Description, &attachment.Path)
		new.Image = attachment
		news = append(news, new)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return news
}