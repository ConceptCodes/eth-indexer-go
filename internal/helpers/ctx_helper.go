package helpers

import (
	"context"
	"net/http"

	"github.com/conceptcodes/eth-indexer-go/internal/constants"
)

type ContextKey string

const (
	RequestIDKey ContextKey = constants.RequestIdCtxKey
	UserId       ContextKey = constants.UserIdCtxKey
	ApiKey       ContextKey = constants.ApiKeyCtxKey
)

func SetRequestId(r *http.Request, requestID string) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, RequestIDKey, requestID)
	return r.WithContext(ctx)
}

func GetRequestId(r *http.Request) string {
	return r.Context().Value(RequestIDKey).(string)
}

func SetUserId(r *http.Request, userId string) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, UserId, userId)
	return r.WithContext(ctx)
}

func GetUserId(r *http.Request) string {
	userId := r.Context().Value(UserId)
	if userId == nil {
		return ""
	}
	return userId.(string)
}

func SetApiKey(r *http.Request, apiKey string) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, ApiKey, apiKey)
	return r.WithContext(ctx)
}

func GetApiKey(r *http.Request) string {
	apiKey := r.Context().Value(ApiKey)
	if apiKey == nil {
		return ""
	}
	return apiKey.(string)
}
