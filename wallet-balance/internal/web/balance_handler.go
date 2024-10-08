package web

import (
	"encoding/json"
	"net/http"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/find_account"
	"github.com/go-chi/chi/v5"
)

type WebBalanceHandler struct {
	FindAccountUseCase find_account.FindAccountUseCase
}

func NewWebBalanceHandler(f find_account.FindAccountUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		FindAccountUseCase: f,
	}
}

func (h *WebBalanceHandler) FindBalance(w http.ResponseWriter, r *http.Request) {
	accountId:= chi.URLParam(r, "account_id")

	input := find_account.FindAccountInputDTO{
		ID: accountId,
	}

	account, err := h.FindAccountUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
