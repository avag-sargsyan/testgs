package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avag-sargsyan/testgs/internal/conf"
	"github.com/avag-sargsyan/testgs/internal/usecase"
)

type PackHandler interface {
	Pack(w http.ResponseWriter, r *http.Request)
}

func NewPackHandler(conf *conf.App) PackHandler {
	return &packHandler{packService: usecase.GetPackService(conf)}
}

type packHandler struct {
	packService usecase.PackService
}

func (h *packHandler) Pack(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query().Get("order")
	if order == "" {
		http.Error(w, "Missing 'order' parameter", http.StatusBadRequest)
		return
	}

	// extract integer from order and store in orderInt
	var orderInt int
	fmt.Sscanf(order, "%d", &orderInt)

	packs := h.packService.CalculatePacks(orderInt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packs)
}
