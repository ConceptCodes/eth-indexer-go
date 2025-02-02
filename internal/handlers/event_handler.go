package handlers

import (
	"fmt"
	"net/http"

	"github.com/conceptcodes/eth-indexer-go/internal/constants"
	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/gorilla/mux"

	"github.com/rs/zerolog"
)

type EventHandler struct {
	log            *zerolog.Logger
	eventLogRepo   repository.EventLogRepository
	responseHelper *helpers.ResponseHelper
}

func NewEventHandler(
	log *zerolog.Logger,
	eventLogRepo repository.EventLogRepository,
	responseHelper *helpers.ResponseHelper,
) *EventHandler {
	return &EventHandler{
		log:            log,
		eventLogRepo:   eventLogRepo,
		responseHelper: responseHelper,
	}
}

func (h *EventHandler) GetEventLogsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		h.responseHelper.SendErrorResponse(w, "Transaction Hash is required", constants.BadRequest, nil)
		return
	}

	eventLogs, err := h.eventLogRepo.FindByTransactionHash(address)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.InternalServerError, err)
		return
	}

	msg := fmt.Sprintf(constants.GetEntityByIdMessage, "Event Logs", address)
	h.responseHelper.SendSuccessResponse(w, msg, models.SimplifyEvents(*eventLogs))
}
