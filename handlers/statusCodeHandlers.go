package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func PostError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func StatusCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["statuscode"]

	switch id {
	case "200":
		PostError(w, http.StatusOK)
		return
	case "400":
		PostError(w, http.StatusBadRequest)
		return
	case "401":
		PostError(w, http.StatusUnauthorized)
		return
	case "403":
		PostError(w, http.StatusForbidden)
		return
	case "404":
		PostError(w, http.StatusNotFound)
		return
	case "405":
		PostError(w, http.StatusMethodNotAllowed)
		return
	case "429":
		PostError(w, http.StatusTooManyRequests)
		return
	case "500":
		PostError(w, http.StatusInternalServerError)
		return
	case "501":
		PostError(w, http.StatusNotImplemented)
		return
	case "502":
		PostError(w, http.StatusBadGateway)
		return
	case "503":
		PostError(w, http.StatusServiceUnavailable)
		return
	case "504":
		PostError(w, http.StatusGatewayTimeout)
		return
	default:
		PostError(w, http.StatusNotFound)
		return
	}
}
