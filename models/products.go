package models

import (
	"bitbucket.org/agoalofalife/shop.game/sessions"
	"fmt"
)

type Products struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Created_at  NullTime
	Update_at   NullTime
	Deleted_at  NullTime
	Image       Attachments
}

type ProductCart struct {
	Products
	Count int
}

func ProductsAll() []Products {
	var (
		err        error
		product    Products
		products   []Products
		attachment Attachments
	)

	results, err := DB.Query("SELECT products.id, products.name, products.description, products.price, attachments.path FROM `products` LEFT JOIN attachments ON products.image = attachments.id")
	defer results.Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &attachment.Path)
		product.Image = attachment
		products = append(products, product)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return products
}

func ProductsByCategory(category int) []Products {
	var (
		err        error
		product    Products
		products   []Products
		attachment Attachments
	)

	results, err := DB.Query("SELECT products.id, products.name, products.description, products.price, attachments.path " +
		"FROM products " +
		"LEFT JOIN attachments ON products.image = attachments.id LEFT JOIN category_product ON products.id = category_product.id " +
		"WHERE category_product.category_id = 1")
	defer results.Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &attachment.Path)
		product.Image = attachment
		products = append(products, product)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return products
}

func ProductListFromCart(products []sessions.Cart) (productCart []ProductCart) {
	var (
		product Products
		attachment Attachments
	)

	for _, item := range products {
		results := DB.QueryRow("SELECT products.id, products.name, products.description, products.price, attachments.path FROM `products` LEFT JOIN attachments ON products.image = attachments.id WHERE products.id = ?", item.Id)
		results.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &attachment.Path)
		productCart = append(productCart, ProductCart{product, item.Count})
	}
	return
}
