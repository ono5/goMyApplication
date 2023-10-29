package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tpml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tpml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
