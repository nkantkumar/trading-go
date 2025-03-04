package api

import (
	"encoding/json"
	"net/http"
	"trading-go/internal/order"
)

type Handler struct {
	orderService *order.Service
}

func NewHandler(service *order.Service) *Handler {
	return &Handler{orderService: service}
}

func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var o order.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.orderService.PlaceOrder(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.ListOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
