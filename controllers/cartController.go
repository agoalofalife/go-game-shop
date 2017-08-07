package controllers

import (
	"bitbucket.org/agoalofalife/shop.game/sessions"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	fmt.Println(r.URL.String())
	sessions.AddProductInSession(r, w, int(id), 1)
	http.Redirect(w, r, r.URL.String(), 302)
}
