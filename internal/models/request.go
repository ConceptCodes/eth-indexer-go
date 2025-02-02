package models

type Request struct {
	ID     string
	ApiKey string
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required,noSQLKeywords"`
	Password string `json:"password" validate:"required,noSQLKeywords"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"email,required,noSQLKeywords"`
	Password string `json:"password" validate:"required,noSQLKeywords"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"email,required,noSQLKeywords"`
}

type ResetPasswordRequest struct {
	Password string `json:"password" validate:"email,required,noSQLKeywords"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"email,required,noSQLKeywords"`
	Otp   string `json:"otp" validate:"required,noSQLKeywords,numeric"`
}
