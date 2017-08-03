package main

import (
	"fmt"
	c "github.com/agoalofalife/shop.game/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/agoalofalife/shop.game/views"
)

func main() {
	r := mux.NewRouter()
	controller := c.New()

	views.LoadTemplates()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`./static/`))))
	r.HandleFunc("/", controller.Index)

	fmt.Println(`Start server ... localhost:8080`)
	log.Fatal(http.ListenAndServe(":8080", r))
}
