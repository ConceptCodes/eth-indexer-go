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

type BlockHandler struct {
	log            *zerolog.Logger
	blockRepo      repository.BlockRepository
	responseHelper *helpers.ResponseHelper
}

func NewBlockHandler(
	log *zerolog.Logger,
	blockRepo repository.BlockRepository,
	responseHelper *helpers.ResponseHelper,
) *BlockHandler {
	return &BlockHandler{
		log:            log,
		blockRepo:      blockRepo,
		responseHelper: responseHelper,
	}
}

func (h *BlockHandler) GetBlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]

	if blockNumber == "" {
		h.responseHelper.SendErrorResponse(w, "Block Number is required", constants.BadRequest, nil)
		return
	}

	block, err := h.blockRepo.FindByBlockNumber(blockNumber)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.InternalServerError, err)
		return
	}

	msg := fmt.Sprintf(constants.GetEntityByIdMessage, "Block", blockNumber)
	h.responseHelper.SendSuccessResponse(w, msg, block.SimpleBlock())
}
