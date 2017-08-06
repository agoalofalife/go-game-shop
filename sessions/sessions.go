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
