package sessions

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

const product = "product"

// structure for save cart in session
type Cart struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

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

	if session.Values[product] != nil {
		var cartSession []Cart
		var flagExistId bool = false
		jsonString := session.Values[product]
		json.Unmarshal([]byte(jsonString.(string)), &cartSession)

		for key, item := range cartSession {
			if item.Id == id {
				cartSession[key].Count = item.Count + 1
				jsonString, _ := json.Marshal(cartSession)
				session.Values[product] = string(jsonString)
				flagExistId = true
			}
		}

		if flagExistId == false {
			c := Cart{id, count}
			cartSession = append(cartSession, c)
			jsonString, _ := json.Marshal(cartSession)
			session.Values[product] = string(jsonString)
		}
	} else {
		c := Cart{id, count}
		jsonString, _ := json.Marshal([]Cart{c})
		session.Values[product] = string(jsonString)
	}
	session.Save(r, w)
}

func CountCartConcreteSession(r *http.Request) (count int) {
	var cartSession []Cart
	var countProduct int
	session, _ := store.Get(r, "session")
	jsonString := session.Values[product]
	json.Unmarshal([]byte(jsonString.(string)), &cartSession)

	for _, product := range cartSession {
		countProduct = countProduct + product.Count
	}

	return countProduct
}

func ListProductFromCart(r *http.Request) (cartSession []Cart) {
	session, _ := store.Get(r, "session")
	jsonString := session.Values[product]
	json.Unmarshal([]byte(jsonString.(string)), &cartSession)
	return
}
