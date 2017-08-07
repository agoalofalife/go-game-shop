package main

import (
	c "bitbucket.org/agoalofalife/shop.game/controllers"
	"bitbucket.org/agoalofalife/shop.game/models"
	"bitbucket.org/agoalofalife/shop.game/views"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	loadEnvironment()
	models.Init()
	views.LoadTemplates()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`./static/`))))
	r.HandleFunc("/", c.BaseIndex)
	r.HandleFunc("/products/category/{id}", c.ProductIndex)
	r.HandleFunc("/about", c.BasePageAbout)
	r.HandleFunc("/news", c.BasePageNews)
	r.HandleFunc("/news/{id}", c.BasePageNewsConcrete)
	r.HandleFunc("/addCart/{id}", c.AddCart)
	r.HandleFunc("/my-orders", c.BasePageMyOrders)
	r.HandleFunc("/registration", c.AuthRegistration).Methods("POST")
	r.HandleFunc("/logout", c.AuthLogout)
	r.HandleFunc("/login", c.AuthLogIn).Methods("POST")

	fmt.Println(`Start server ... localhost:8080`)
	log.Fatal(http.ListenAndServe(":8080", r))
}
