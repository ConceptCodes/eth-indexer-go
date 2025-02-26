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

func (h *ViewHandler) GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	txCount := int64(0)
	blockCount := int64(0)
	avgBlockTime := uint64(0)

	txCount, _ = h.transactionRepo.Count()
	blockCount, _ = h.blockRepo.Count()

	ptrTransactions, _ := h.transactionRepo.Recent(10)
	transactions := make([]models.Transaction, len(ptrTransactions))
	for i, t := range ptrTransactions {
		transactions[i] = *t
	}
	// blocks, _ := h.blockRepo.GetAll()
	// if len(blocks) > 1 {
	// 	diff := blocks[0].Timestamp - blocks[len(blocks)-1].Timestamp
	// 	avgBlockTime = diff / uint64(len(blocks)-1) / 1000000000 / 60
	// }

	data := models.HomeData{
		TxCount:      txCount,
		BlockCount:   blockCount,
		Transactions: models.SimplifyTransactions(transactions),
		AvgBlockTime: avgBlockTime,
	}

	views.Index(data).Render(r.Context(), w)
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
	if err != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}

	txs := make([]models.Transaction, len(transactions))
	for i, tx := range transactions {
		txs[i] = *tx
	}
	simpleTxs := models.SimplifyTransactions(txs)
	totalPages := len(txs) / sizeInt
	if len(txs)%sizeInt != 0 {
		totalPages++
	}

	data := models.BlockData{
		Block:      block.SimpleBlock(),
		Txs:        simpleTxs,
		PageNumber: int64(pageInt),
		TotalPages: int64(totalPages),
		TxCount:    int64(len(txs)),
	}

	views.Block(data).Render(r.Context(), w)
}

func (h *ViewHandler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	query := r.URL.Query()

	pageNumber := query.Get("page")
	pageSize := query.Get("size")
	txCount := int64(0)

	if address == "" {
		views.NotFound().Render(r.Context(), w)
		return
	}

	if pageNumber == "" {
		pageNumber = "1"
	}

	if pageSize == "" {
		pageSize = strconv.FormatInt(constants.DefaultPageSize, 10)
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

	transactions, err := h.transactionRepo.FindByFromAccount(address, pageInt, sizeInt)
	if err != nil {
		views.NotFound().Render(r.Context(), w)
		return
	}

	txs := make([]models.Transaction, len(transactions))
	for i, tx := range transactions {
		txs[i] = *tx
	}

	txCount, _ = h.transactionRepo.Count()

	data := models.AccountData{
		Address:    address,
		Txs:        models.SimplifyTransactions(txs),
		TxCount:    txCount,
		PageSize:   int64(sizeInt),
		PageNumber: int64(pageInt),
		TotalPages: txCount / int64(sizeInt),
	}

	views.Account(data).Render(r.Context(), w)
}

func (h *ViewHandler) Get404Handler(w http.ResponseWriter, r *http.Request) {
	views.NotFound().Render(r.Context(), w)
}
