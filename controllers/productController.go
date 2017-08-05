package controllers

import (
	"bitbucket.org/agoalofalife/shop.game/models"
	"bitbucket.org/agoalofalife/shop.game/views"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	m := map[string]interface{}{}
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = models.ProductsByCategory(int(id))
	m["categoryName"] = models.CategoryName(int(id)).Name
	m["categories"] = models.CategoriesAll()
	m["news"] = models.NewsAll()
	views.ProductsCategoryTemplate.Execute(w, m)

}
