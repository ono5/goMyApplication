package main

import (
	"fmt"
	"net/http"

	"github.com/ono5/myGoWebApplication/pkg/handler"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
