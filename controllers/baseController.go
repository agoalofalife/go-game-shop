package controllers


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/agoalofalife/shop.game/views"
	"net/http"
	"github.com/agoalofalife/shop.game/models"
)

type baseController struct {}
type WebData struct {
	Title string
}
func (c *baseController) Index(w http.ResponseWriter, r *http.Request)  {
	m := map[string]interface{}{}
	m["categories"] = models.CategoriesAll()
	m["products"] = models.ProductsAll()

	views.HomeTemplate.Execute(w, m)

}

func New() *baseController  {
	c := &baseController{}
	return c
}
