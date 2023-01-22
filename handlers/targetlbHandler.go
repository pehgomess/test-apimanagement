package handlers

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Server string `json:"server"`
}
type Servers []Server

var server = "server1"

func ServerName() string {
	if len(os.Args) != 1 {
		switch os.Args[1] {
		case "server2":
			server := "server2"
			return server
		default:
			server := "server1"
			return server
		}
	}
	return server
}

func Tl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	servers := Server{
		Server: ServerName(),
	}

	file, err := os.Open("/tmp/goerr")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Text() == "yes" {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(servers)
}
