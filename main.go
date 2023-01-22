package main

import (
	"cluster1/handlers"
	"fmt"
	"net/http"
	"os"
)

func handlerRequest() {
	http.HandleFunc("/tl", handlers.Tl)
	http.HandleFunc("/code/", handlers.StatusCode)
	http.HandleFunc("/status", handlers.StatusUP)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func main() {
	handlers.ServerName()
	handlerRequest()
}
