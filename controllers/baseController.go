package controllers

import (
	"bitbucket.org/agoalofalife/shop.game/models"
	"bitbucket.org/agoalofalife/shop.game/sessions"
	"bitbucket.org/agoalofalife/shop.game/views"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func BaseIndex(w http.ResponseWriter, r *http.Request) {

	m := map[string]interface{}{}

	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = products
	m["categories"] = models.CategoriesAll()
	m["news"] = models.NewsAll()
	m["existUser"] = sessions.IsLoggedIn(r)
	views.HomeTemplate.Execute(w, m)
}

func BasePageAbout(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}

	m["categories"] = models.CategoriesAll()
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = products
	m["news"] = models.NewsAll()
	m["existUser"] = sessions.IsLoggedIn(r)
	views.PageAbout.Execute(w, m)

}

func BasePageNews(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}

	m["categories"] = models.CategoriesAll()
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["news"] = models.NewsAll()
	m["existUser"] = sessions.IsLoggedIn(r)
	views.News.Execute(w, m)
}

func BasePageNewsConcrete(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)

	m["categories"] = models.CategoriesAll()
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = products[:3]
	m["new"] = models.NewById(int(id))
	m["existUser"] = sessions.IsLoggedIn(r)
	views.New.Execute(w, m)
}

func BasePageMyOrders(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}
	m["categories"] = models.CategoriesAll()
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["news"] = models.NewsAll()
	m["existUser"] = sessions.IsLoggedIn(r)
	views.MyOrders.Execute(w, m)
}
