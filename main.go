package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

// Home page
func Home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "home.page.tmpl")
}

// About page
func About(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "about.page.tmpl")
}

func renderTemplate(res http.ResponseWriter, tmpl string) {
	parsedTemplate, errParseFiles := template.ParseFiles("./templates/" + tmpl)
	if errParseFiles != nil {
		fmt.Println("error parsing files:", errParseFiles)
		return
	}

	errTemplateExecute := parsedTemplate.Execute(res, nil)
	if errTemplateExecute != nil {
		fmt.Println("error parsing template:", errTemplateExecute)
		return
	}
}

func checkErr(text string, err error) {
	if err != nil {
		fmt.Println(text, ":", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("server is running")
	_ = http.ListenAndServe(port, nil)

}
