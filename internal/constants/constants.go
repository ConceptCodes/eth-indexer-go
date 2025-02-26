package constants

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"

	TraceIdHeader = "x-request-id"
	ApiKeyHeader  = "x-api-key"

	RequestIdCtxKey = "request_id"
	ApiKeyCtxKey    = "api_key"
	UserIdCtxKey    = "user_id"

	NotFound            = "eth-indx-404"
	BadRequest          = "eth-indx-400"
	Unauthorized        = "eth-indx-401"
	Forbidden           = "eth-indx-403"
	InternalServerError = "eth-indx-500"

	DefaultRedisTtl = 10 * time.Minute
	DefaultPageSize = 25

	EntityNotFound             = "%s with %s %s does not exist."
	GetEntityByIdMessage       = "Found %s with id %s."
	SaveEntityError            = "Error while saving %s."
	SuccessMessage             = "You have successfully %s!"
	CreateEntityError          = "Error while creating %s."
	CreateEntityMessage        = "Created %s successfully."
	OtpCodeMessage             = "Your OTP code is %s."
	InternalServerErrorMessage = "Internal server error."

	IndexViewEndpoint       = "/"
	LoginViewEndpoint       = "/login"
	RegisterViewEndpoint    = "/register"
	TransactionViewEndpoint = "/tx/{hash:0x[0-9a-fA-F]{64}}"
	AccountViewEndpoint     = "/account/{address:0x[0-9a-fA-F]{40}}"
	BlockViewEndpoint       = "/block/{blockNumber:[0-9]+}"
	NotFoundViewEndpoint    = "/404"

	ApiPrefix           = "/api/v1"
	HealthCheckEndpoint = ApiPrefix + "/health/alive"
	ReadinessEndpoint   = ApiPrefix + "/health/status"

	LoginEndpoint          = ApiPrefix + "/auth/login"
	RegisterEndpoint       = ApiPrefix + "/auth/register"
	ForgotPasswordEndpoint = ApiPrefix + "/auth/forgot-password"
	VerifyEmailEndpoint    = ApiPrefix + "/auth/verify-email"
	ResetPasswordEndpoint  = ApiPrefix + "/auth/reset-password"

	GetTransactionsByHashEndpoint      = ApiPrefix + TransactionViewEndpoint
	GetBlockByNumberEndpoint           = ApiPrefix + BlockViewEndpoint
	GetEventsByContractAddressEndpoint = ApiPrefix + "/events/{address:0x[0-9a-fA-F]{40}}"
)
