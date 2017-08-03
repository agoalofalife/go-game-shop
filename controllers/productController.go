package controllers

import (
	"net/http"
	"fmt"
	"os"
	"github.com/gorilla/mux"
	"github.com/agoalofalife/shop.game/models"
	"strconv"
	"github.com/agoalofalife/shop.game/views"
)



func  ProductIndex(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	m := map[string]interface{}{}
	m["products"] = models.ProductsByCategory(int(id))
	//products := models.ProductsAll()
	//m["randomProduct"] = products[(rand.Int() % len(products))]
	//m["products"] = products
	//
	views.HomeTemplate.Execute(w, m)
	fmt.Println(m["products"])
	os.Exit(0)
}