package main

import "net/http"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home-page.tpml")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about-page.tpml")
}
