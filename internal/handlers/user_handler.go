package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/conceptcodes/eth-indexer-go/internal/constants"
	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/conceptcodes/eth-indexer-go/pkg/email"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	userRepo        repository.UserRepository
	authRepo        repository.AuthRepository
	log             *zerolog.Logger
	authHelper      *helpers.AuthHelper
	responseHelper  *helpers.ResponseHelper
	validatorHelper *helpers.ValidatorHelper
	emailClient     *email.EmailClient
	redisHelper     *helpers.RedisHelper
}

func NewUserHandler(
	userRepo repository.UserRepository,
	authRepo repository.AuthRepository,
	log *zerolog.Logger,
	authHelper *helpers.AuthHelper,
	responseHelper *helpers.ResponseHelper,
	validatorHelper *helpers.ValidatorHelper,
	emailClient *email.EmailClient,
	redisHelper *helpers.RedisHelper,
) *UserHandler {
	return &UserHandler{
		userRepo:        userRepo,
		authRepo:        authRepo,
		log:             log,
		authHelper:      authHelper,
		responseHelper:  responseHelper,
		validatorHelper: validatorHelper,
		emailClient:     emailClient,
		redisHelper:     redisHelper,
	}
}

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var data models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.BadRequest, err)
	}

	h.validatorHelper.ValidateStruct(w, &data)

	password_hash, err := h.authHelper.HashPassword(data.Password)
	err_message := fmt.Sprintf(constants.CreateEntityError, "User")

	if err != nil {
		h.log.Error().Err(err).Msg("Error hashing password")
		h.responseHelper.SendErrorResponse(w, err_message, constants.InternalServerError, err)
	}

	apiKey := uuid.New().String()

	user := models.User{
		Email:    data.Email,
		Password: password_hash,
		Name:     data.Name,
		Enabled:  false,
		ApiKey:   apiKey,
	}

	err = h.userRepo.Create(&user)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err_message, constants.InternalServerError, err)
	}

	code, err := h.authHelper.GenerateOtpCode(data.Email)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, "Error w/ otp generation", constants.InternalServerError, err)
	}

	msg := fmt.Sprintf(constants.OtpCodeMessage, code)
	h.redisHelper.SetData(data.Email, code, constants.DefaultRedisTtl)
	err = h.emailClient.SendEmail(data.Email, "verify-email", msg)

	if err != nil {
		h.log.Error().Err(err).Msg("Error sending verification email")
		h.responseHelper.SendErrorResponse(w, "Error w/ sending verification email", constants.InternalServerError, err)
	}

	res := &models.RegisterUserResponse{
		Email:  data.Email,
		Name:   data.Name,
		ApiKey: apiKey,
	}

	h.responseHelper.SendSuccessResponse(w, "User registered successfully", res)

}

func (h *UserHandler) VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	var data models.VerifyEmailRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.BadRequest, err)
	}

	h.validatorHelper.ValidateStruct(w, &data)

	err = h.authHelper.ValidateOtpCode(data.Email, data.Otp)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, "Error verifying OTP code", constants.InternalServerError, err)
	}

	user, err := h.userRepo.FindByEmail(data.Email)

	if err != nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User ", "email:", data.Email)
		h.responseHelper.SendErrorResponse(w, err_message, constants.InternalServerError, err)
	}

	if user == nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User", "email: ", data.Email)
		h.responseHelper.SendErrorResponse(w, err_message, constants.NotFound, nil)
	} else {
		user.EmailVerified = time.Now()
		user.Enabled = true
		err = h.userRepo.Save(user)

		if err != nil {
			h.responseHelper.SendErrorResponse(w, "Error verifying email", constants.InternalServerError, err)
		}
	}

	h.responseHelper.SendSuccessResponse(w, "Email verified successfully", nil)

}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var data models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.BadRequest, err)
	}

	h.validatorHelper.ValidateStruct(w, &data)

	user, err := h.userRepo.FindByEmail(data.Email)

	if err != nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User ", "email:", data.Email)
		h.responseHelper.SendErrorResponse(w, err_message, constants.InternalServerError, err)
	}

	if user == nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User", "email: ", data.Email)
		h.responseHelper.SendErrorResponse(w, err_message, constants.NotFound, nil)
	} else {

		if user.EmailVerified.After(time.Now()) {
			h.responseHelper.SendErrorResponse(w, "Email not verified", constants.BadRequest, err)
		}
	}

	valid := h.authHelper.CheckPasswordHash(data.Password, user.Password)

	if !valid {
		h.responseHelper.SendErrorResponse(w, "Invalid credentials", constants.BadRequest, err)
	}

	h.responseHelper.SendSuccessResponse(w, "Successful login", nil)
}

func (h *UserHandler) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var data models.ForgotPasswordRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.BadRequest, err)
	}

	h.validatorHelper.ValidateStruct(w, &data)

	user, err := h.userRepo.FindByEmail(data.Email)

	if err != nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User ", "email:", data.Email)
		h.responseHelper.SendErrorResponse(w, err_message, constants.InternalServerError, err)
	}

	if user == nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User", "email: ", data.Email)
		h.responseHelper.SendErrorResponse(w, err_message, constants.NotFound, nil)
	} else {
		if user.EmailVerified.After(time.Now()) {
			h.responseHelper.SendErrorResponse(w, "Email not verified", constants.BadRequest, err)
		}
	}

	reset_token := uuid.New().String()

	tmp := models.Auth{
		UserID: user.ID,
		Token:  reset_token,
	}

	err = h.authRepo.Create(&tmp)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, "Error sending reset password email", constants.InternalServerError, err)
	}

	url := fmt.Sprintf("%s?token=%s", r.URL, reset_token)

	msg := fmt.Sprintf("Please click this link to reset your password: %s", url)
	err = h.emailClient.SendEmail(data.Email, "reset-password", msg)

	if err != nil {
		h.log.Error().Err(err).Msg("Error sending reset password email")
		h.responseHelper.SendErrorResponse(w, "Error sending reset password email", constants.InternalServerError, err)
	}

	h.responseHelper.SendSuccessResponse(w, "Reset password email sent successfully", nil)
}

func (h *UserHandler) ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	token := params.Get("token")

	if token == "" {
		h.log.Error().Msg("Token is empty")
		h.responseHelper.SendErrorResponse(w, "Token is empty", constants.InternalServerError, nil)
	}

	var data models.ResetPasswordRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		h.responseHelper.SendErrorResponse(w, err.Error(), constants.BadRequest, err)
	}

	record, err := h.authRepo.FindByToken(token)

	if err == nil {
		err_message := fmt.Sprintf(constants.EntityNotFound, "User ", "email:", "")
		h.responseHelper.SendErrorResponse(w, err_message, constants.BadRequest, err)
	}

	if token == record.Token {
		user, err := h.userRepo.FindByEmail(record.Email)

		if err != nil {
			err_message := fmt.Sprintf(constants.EntityNotFound, "User ", "id:", strconv.Itoa(int(record.UserID)))
			h.responseHelper.SendErrorResponse(w, err_message, constants.BadRequest, err)
		}

		password_hash, err := h.authHelper.HashPassword(data.Password)

		if err != nil {
			h.log.Error().Err(err).Msg("Error hashing password")
			h.responseHelper.SendErrorResponse(w, "Error resetting password", constants.InternalServerError, err)
		}

		user.Password = password_hash

		err = h.userRepo.Save(user)

		if err != nil {
			h.responseHelper.SendErrorResponse(w, "Error resetting password", constants.InternalServerError, err)
		}

	} else {
		h.responseHelper.SendErrorResponse(w, "Invalid token", constants.BadRequest, nil)
	}

	h.responseHelper.SendSuccessResponse(w, "Password reset successfully", nil)

}
