package handlers

import (
	"net/http"
)

func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func StatusCode(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/code/200":
		postError(w, http.StatusOK)
		return
	case "/code/400":
		postError(w, http.StatusBadRequest)
		return
	case "/code/401":
		postError(w, http.StatusUnauthorized)
		return
	case "/code/403":
		postError(w, http.StatusForbidden)
		return
	case "/code/404":
		postError(w, http.StatusNotFound)
		return
	case "/code/405":
		postError(w, http.StatusMethodNotAllowed)
		return
	case "/code/429":
		postError(w, http.StatusTooManyRequests)
		return
	case "/code/500":
		postError(w, http.StatusInternalServerError)
		return
	case "/code/501":
		postError(w, http.StatusNotImplemented)
		return
	case "/code/502":
		postError(w, http.StatusBadGateway)
		return
	case "/code/503":
		postError(w, http.StatusServiceUnavailable)
		return
	case "/code/504":
		postError(w, http.StatusGatewayTimeout)
		return
	default:
		postError(w, http.StatusNotFound)
		return
	}
}
