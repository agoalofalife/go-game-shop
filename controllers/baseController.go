package controllers


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/agoalofalife/shop.game/views"
	"net/http"
	"github.com/agoalofalife/shop.game/models"
	"math/rand"
)



func BaseIndex(w http.ResponseWriter, r *http.Request)  {

	m := map[string]interface{}{}

	m["categories"] = models.CategoriesAll()
	products := models.ProductsAll()
	m["randomProduct"] = products[(rand.Int() % len(products))]
	m["products"] = products

	views.HomeTemplate.Execute(w, m)

}
