package controllers

import (
	"bitbucket.org/agoalofalife/shop.game/sessions"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	sessions.AddProductInSession(r, w, int(id), 1)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
