package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

const product = "product"

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

func RemoveSession(r *http.Request, w http.ResponseWriter) {
	session, _ := store.Get(r, "session")
	session.Values["loggedIn"] = "false"
	session.Values["id"] = ""
	session.Save(r, w)
}

func AddProductInSession(r *http.Request, w http.ResponseWriter, id int, count int) {
	session, _ := store.Get(r, "session")
	if session.Values[product] == nil {
		session.Values[product] = map[int]int{}
	}

	session.Values[product].(map[int]int)[id] = count
	session.Save(r, w)
}

func CountCartConcreteSession(r *http.Request) (count int) {
	session, _ := store.Get(r, "session")
	if session.Values[product] == nil {
		session.Values[product] = map[int]int{}
	}
	return len(session.Values[product].(map[int]int))
}
