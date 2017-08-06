package controllers

import (
	"bitbucket.org/agoalofalife/shop.game/models"
	"bitbucket.org/agoalofalife/shop.game/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func AuthRegistration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := map[string]interface{}{}
	m["password"], _ = bcrypt.GenerateFromPassword([]byte(r.Form.Get(`password`)), bcrypt.DefaultCost)
	m["name"] = r.Form.Get(`name`)
	m["email"] = r.Form.Get(`email`)
	id := models.UserCreate(m)
	sessions.AddSession(id, r, w)
	http.Redirect(w, r, "/", 302)
}

func AuthLogout(w http.ResponseWriter, r *http.Request) {
	sessions.RemoveSession(r, w)
	http.Redirect(w, r, "/", 302)
}

func AuthLogIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := models.UserGetByEmail(r.Form.Get(`email`))

	if user.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form.Get(`password`)))
		if err == nil {
			sessions.AddSession(user.Id, r, w)
			http.Redirect(w, r, "/", 302)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
