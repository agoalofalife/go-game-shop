package views

import (
	"io/ioutil"
	"os"
	"strings"
	"html/template"
	"fmt"
)

var (
	HomeTemplate      *template.Template
	ProductsCategoryTemplate *template.Template
	PageAbout *template.Template
	message           string
	//message will store the message to be shown as notification
	err               error
)

//PopulateTemplates is used to parse all templates present in
//the templates folder
func LoadTemplates() {
	var allFiles []string
	templatesDir := "./views/template/"

	files, err := ioutil.ReadDir(templatesDir)

	if err != nil {
		fmt.Println("Error reading template dir")
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	templates := template.Must(template.ParseFiles(allFiles...))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	HomeTemplate = templates.Lookup("index.html")
	ProductsCategoryTemplate = templates.Lookup("category.html")
	PageAbout = templates.Lookup("about.html")
}