package main

import (
	"fmt"
	c "github.com/agoalofalife/shop.game/controllers"
	"github.com/agoalofalife/shop.game/models"
	"github.com/agoalofalife/shop.game/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	controller := c.New()

	loadEnvironment()
	models.Init()
	views.LoadTemplates()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`./static/`))))
	r.HandleFunc("/", controller.Index)

	fmt.Println(`Start server ... localhost:8080`)
	log.Fatal(http.ListenAndServe(":8080", r))
}
