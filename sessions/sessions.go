package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func IsLoggedIn(r *http.Request) bool {
	session, _ := store.Get(r, "session")
	if session.Values["loggedIn"] == "true" {
		return true
	}
	return false
}

func AddSession(id int, r *http.Request, w http.ResponseWriter) {
	session, _ := store.Get(r, "session")
	session.Values["loggedIn"] = "true"
	session.Values["id"] = id
	session.Save(r, w)
}
