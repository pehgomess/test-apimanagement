package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func StatusCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["statuscode"]

	switch id {
	case "200":
		postError(w, http.StatusOK)
		return
	case "400":
		postError(w, http.StatusBadRequest)
		return
	case "401":
		postError(w, http.StatusUnauthorized)
		return
	case "403":
		postError(w, http.StatusForbidden)
		return
	case "404":
		postError(w, http.StatusNotFound)
		return
	case "405":
		postError(w, http.StatusMethodNotAllowed)
		return
	case "429":
		postError(w, http.StatusTooManyRequests)
		return
	case "500":
		postError(w, http.StatusInternalServerError)
		return
	case "501":
		postError(w, http.StatusNotImplemented)
		return
	case "502":
		postError(w, http.StatusBadGateway)
		return
	case "503":
		postError(w, http.StatusServiceUnavailable)
		return
	case "504":
		postError(w, http.StatusGatewayTimeout)
		return
	default:
		postError(w, http.StatusNotFound)
		return
	}
}
