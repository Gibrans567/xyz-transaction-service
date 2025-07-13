package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xyz-multifinance/transaction-service/internal/domain"
	"xyz-multifinance/transaction-service/internal/usecase"
)

type TransactionHandler struct {
	uc *usecase.TransactionUsecase
}

func NewTransactionHandler(uc *usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{uc: uc}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, `{"error": "Request body tidak valid"}`, http.StatusBadRequest)
		return
	}

	customerIDFromToken, ok := r.Context().Value(contextKeyCustomerID).(int64)
	if !ok {
		http.Error(w, `{"error": "Gagal mendapatkan ID konsumen dari token"}`, http.StatusUnauthorized)
		return
	}
	tx.CustomerID = customerIDFromToken

	err := h.uc.CreateTransaction(r.Context(), &tx)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusUnprocessableEntity)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transaksi berhasil dibuat"})
}