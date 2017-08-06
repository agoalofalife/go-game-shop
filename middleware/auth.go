package middleware

import (
	"bitbucket.org/agoalofalife/shop.game/sessions"
	"net/http"
)

func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(r) {
			http.Redirect(w, r, "/login/", 302)
			return
		}
		handler(w, r)
	}
}
