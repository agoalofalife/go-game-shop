package views

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	HomeTemplate             *template.Template
	ProductsCategoryTemplate *template.Template
	PageAbout                *template.Template
	News                     *template.Template
	New                      *template.Template
)

func LoadTemplates() {
	var allFiles []string
	templatesDir := "./views/template/"

	funcMap := template.FuncMap{
		"TimeStamp": func(time time.Time) string {
			return time.Format("2006-01-02")
		},
		"Eq": func(a, b interface{}) bool {
			return a == b
		},
		"TrimSpace": func(time time.Time) string {
			return strings.TrimSpace(time.String())
		},
	}
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

	templates := template.Must(template.New("myTemplate").Funcs(funcMap).ParseFiles(allFiles...))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	HomeTemplate = templates.Lookup("index.html")
	ProductsCategoryTemplate = templates.Lookup("category.html")
	PageAbout = templates.Lookup("about.html")
	News = templates.Lookup("news.html")
	New = templates.Lookup("new.html")

}
