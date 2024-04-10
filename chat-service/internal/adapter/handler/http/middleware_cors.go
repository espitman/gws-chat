package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func corsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Allow-Methods", "*")
		next(w, r, ps)
	}
}
