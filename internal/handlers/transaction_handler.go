package handlers

import (
	"fmt"
	"net/http"

	"github.com/conceptcodes/eth-indexer-go/internal/constants"
	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/gorilla/mux"

	"github.com/rs/zerolog"
)

type TransactionHandler struct {
	transactionRepo repository.TransactionRepository
	log             *zerolog.Logger
	responseHelper  *helpers.ResponseHelper
}

func NewTransactionHandler(
	transactionRepo repository.TransactionRepository,
	log *zerolog.Logger,
	responseHelper *helpers.ResponseHelper,
) *TransactionHandler {
	return &TransactionHandler{
		transactionRepo: transactionRepo,
		log:             log,
		responseHelper:  responseHelper,
	}
}

func (h *TransactionHandler) GetTransactionByHashHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	if hash == "" {
		h.responseHelper.SendErrorResponse(w, "Hash is required", constants.BadRequest, nil)
		return
	}

	transaction, err := h.transactionRepo.FindByHash(hash)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.InternalServerError, err)
		return
	}

	msg := fmt.Sprintf(constants.GetEntityByIdMessage, "Transaction", hash)
	h.responseHelper.SendSuccessResponse(w, msg, transaction.SimpleTransaction())

}
