package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "Welcome!\n")
}
