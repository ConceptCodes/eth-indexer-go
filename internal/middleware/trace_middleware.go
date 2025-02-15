package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/conceptcodes/eth-indexer-go/internal/constants"
	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
)

type TraceRequestMiddleware struct {
	log            *zerolog.Logger
	authHelper     *helpers.AuthHelper
	responseHelper *helpers.ResponseHelper
}

func NewTraceRequestMiddleware(log *zerolog.Logger, authHelper *helpers.AuthHelper, responseHelper *helpers.ResponseHelper) *TraceRequestMiddleware {
	return &TraceRequestMiddleware{
		log:            log,
		authHelper:     authHelper,
		responseHelper: responseHelper,
	}
}

func (m *TraceRequestMiddleware) Start(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestId := r.Header.Get(constants.TraceIdHeader)
		if requestId == "" {
			requestId = uuid.New().String()
		}

		w.Header().Add(constants.TraceIdHeader, requestId)

		// apiKey := r.Header.Get(constants.ApiKeyHeader)

		// if apiKey == "" {
		// 	m.responseHelper.SendErrorResponse(w, "API key is required", constants.Unauthorized, nil)
		// 	return
		// }

		// if apiKey != "" {

		// 	valid := m.authHelper.ValidateApiKey(apiKey)

		// 	if !valid {
		// 		m.responseHelper.SendErrorResponse(w, "API key is required", constants.Unauthorized, nil)
		// 		return
		// 	}

		// 	r = helpers.SetApiKey(r, apiKey)
		// }

		r = helpers.SetRequestId(r, requestId)

		next.ServeHTTP(w, r)
	})
}
