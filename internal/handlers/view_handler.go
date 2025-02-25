package handlers

import (
	"net/http"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/conceptcodes/eth-indexer-go/views"
	"github.com/gorilla/mux"
)

type ViewHandler struct {
	transactionRepo repository.TransactionRepository
	blockRepo       repository.BlockRepository
}

func NewViewHandler(
	transactionRepo repository.TransactionRepository,
	blockRepo repository.BlockRepository,
) *ViewHandler {
	return &ViewHandler{
		transactionRepo: transactionRepo,
		blockRepo:       blockRepo,
	}
}

func (h *ViewHandler) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}

func (h *ViewHandler) GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	txCount := int64(0)
	blockCount := int64(0)
	transactions := make([]models.Transaction, 0)

	txCount, _ = h.transactionRepo.Count()
	blockCount, _ = h.blockRepo.Count()

	ptrTransactions, _ := h.transactionRepo.Recent(10)
	transactions = make([]models.Transaction, len(ptrTransactions))
	for i, t := range ptrTransactions {
		transactions[i] = *t
	}

	data := models.HomeData{
		TxCount:      txCount,
		BlockCount:   blockCount,
		Transactions: models.SimplifyTransactions(transactions),
	}

	views.Home(data).Render(r.Context(), w)
}

func (h *ViewHandler) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	
	if hash == "" {
		views.NotFound().Render(r.Context(), w)
		return
	}

	tx, err := h.transactionRepo.FindByHash(hash)
	if err != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}

	views.Transaction(tx).Render(r.Context(), w)
}

func (h *ViewHandler) Get404Handler(w http.ResponseWriter, r *http.Request) {
	views.NotFound().Render(r.Context(), w)
}