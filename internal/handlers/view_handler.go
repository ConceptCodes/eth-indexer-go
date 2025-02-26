package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/conceptcodes/eth-indexer-go/internal/constants"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/conceptcodes/eth-indexer-go/views"
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
	avgBlockTime := uint64(0)

	txCount, _ = h.transactionRepo.Count()
	blockCount, _ = h.blockRepo.Count()

	ptrTransactions, _ := h.transactionRepo.Recent(10)
	transactions = make([]models.Transaction, len(ptrTransactions))
	for i, t := range ptrTransactions {
		transactions[i] = *t
	}

	blocks, _ := h.blockRepo.GetAll()
	if len(blocks) > 0 {
		avgBlockTime = blocks[0].Timestamp - blocks[len(blocks)-1].Timestamp
	}


	data := models.HomeData{
		TxCount:      txCount,
		BlockCount:   blockCount,
		Transactions: models.SimplifyTransactions(transactions),
		AvgBlockTime: avgBlockTime,
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

func (h *ViewHandler) GetBlockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]

	query := r.URL.Query()

	pageNumber := query.Get("page")
	pageSize := query.Get("size")

	if pageNumber == "" {
		pageNumber = "1"
	}

	if pageSize == "" {
		pageSize = strconv.FormatInt(constants.DefaultPageSize, 10)
	}

	if blockNumber == "" {
		views.NotFound().Render(r.Context(), w)
		return
	}

	block, err := h.blockRepo.FindByBlockNumber(blockNumber)
	if err != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}

	blockNum, errConv := strconv.ParseUint(blockNumber, 10, 64)
	if errConv != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}
	pageInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}
	sizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}
	transactions, err := h.transactionRepo.FindByBlockNumber(blockNum, pageInt, sizeInt)

	views.Block(block.SimpleBlock(), transactions).Render(r.Context(), w)
}

func (h *ViewHandler) Get404Handler(w http.ResponseWriter, r *http.Request) {
	views.NotFound().Render(r.Context(), w)
}
