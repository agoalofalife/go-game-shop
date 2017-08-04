package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/agoalofalife/shop.game/models"
	"strconv"
	"github.com/agoalofalife/shop.game/views"
	"math/rand"
)



func  ProductIndex(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	m := map[string]interface{}{}
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = models.ProductsByCategory(int(id))
	m["categoryName"] = models.CategoryName(int(id)).Name
	m["categories"] = models.CategoriesAll()
	views.ProductsCategoryTemplate.Execute(w, m)

}