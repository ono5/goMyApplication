package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// renderTemplate serves as a wrapper and renders
// a layout and a template from folder /templates to a desired writer
// func RenderTemplate(w http.ResponseWriter, tpml string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tpml, "./templates/base-layout.tpml")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 	}
// }

// Template Cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the themplate in our chache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("createing template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base-layout.tpml",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
