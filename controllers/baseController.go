package controllers

import (
	"github.com/agoalofalife/shop.game/models"
	"github.com/agoalofalife/shop.game/views"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
)

func BaseIndex(w http.ResponseWriter, r *http.Request) {

	m := map[string]interface{}{}

	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = products
	m["categories"] = models.CategoriesAll()
	m["news"] = models.NewsAll()

	views.HomeTemplate.Execute(w, m)
}

func BasePageAbout(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}

	m["categories"] = models.CategoriesAll()
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = products
	m["news"] = models.NewsAll()

	views.PageAbout.Execute(w, m)

}
