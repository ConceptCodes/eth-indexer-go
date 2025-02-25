package handlers

import (
	"net/http"

	"github.com/conceptcodes/eth-indexer-go/views"
)

type ViewHandler struct {
}

func NewViewHandler() *ViewHandler {
	return &ViewHandler{}
}

func (h *ViewHandler) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}
