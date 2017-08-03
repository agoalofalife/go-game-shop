package controller


import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/agoalofalife/shop.game/views"
	"net/http"
	"fmt"
	"os"
)


type baseController struct {}
type WebData struct {
	Title string
}
func (c *baseController) Index(w http.ResponseWriter, r *http.Request)  {
	// Create an sql.DB and check for errors
	db, err := sql.Open("mysql", "root:test@/shop")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	var databaseUsername  string
	var databasePassword  string
	db.QueryRow("SELECT username, password FROM users").Scan(&databaseUsername, &databasePassword)
	fmt.Println(databaseUsername)
	os.Exit(0)
	// If not then redirect to the login page
	s := WebData{
		"One one"}
	views.HomeTemplate.Execute(w, &s)

}

func New() *baseController  {
	c := &baseController{}
	return c
}
