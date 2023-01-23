package handlers

import (
	"encoding/json"
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
	defer file.Close()
	if err != nil {
		PostError(w, http.StatusServiceUnavailable)
	} else {
		json.NewEncoder(w).Encode(servers)
	}

}
