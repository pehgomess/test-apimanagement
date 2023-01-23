package main

import (
	"cluster1/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
)

var tracer trace.Tracer

func handlerRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/tl", handlers.Tl)
	router.HandleFunc("/code/{statuscode:[0-9]+}", handlers.StatusCode)
	router.HandleFunc("/health", handlers.StatusUP)
	//http.ListenAndServe(":8080", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func main() {
	handlers.ServerName()
	handlerRequest()
}
