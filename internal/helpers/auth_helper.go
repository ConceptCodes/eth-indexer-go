package helpers

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type AuthHelper struct {
	log         *zerolog.Logger
	userRepo    repository.UserRepository
	redisHelper *RedisHelper
	cfg         *config.Config
}

func NewAuthHelper(log *zerolog.Logger, userRepo repository.UserRepository) *AuthHelper {
	return &AuthHelper{log: log, userRepo: userRepo}
}

func (h *AuthHelper) ValidateApiKey(apiKey string) bool {
	h.log.Debug().Msgf("Validating API key: %s", apiKey)

	_, err := h.userRepo.FindByApiKey(apiKey)

	if err != nil {
		h.log.Error().Err(err).Msg("Error validating API key")
		return false
	}

	return true
}

func (h *AuthHelper) GenerateOtpCode(target string) (string, error) {
	otp_code := strconv.Itoa(rand.Intn(9000) + 1000)

	key := fmt.Sprintf("otp:%s", target)
	dur := time.Duration(h.cfg.OtpExpireInMins) * time.Minute

	err := h.redisHelper.SetData(key, otp_code, dur)

	if err != nil {
		h.log.Error().Err(err).Msg("Error generating OTP code")
		return "", err
	}

	return otp_code, nil
}

func (h *AuthHelper) ValidateOtpCode(target string, otpCode string) error {
	key := fmt.Sprintf("otp:%s", target)

	code, err := h.redisHelper.GetData(key)

	if err != nil {
		h.log.Error().Err(err).Msg("Error getting OTP code")
		return err
	}

	if code == "" {
		return errors.New("OTP code not found")
	}

	if otpCode != code {
		return errors.New("invalid OTP code")
	}

	return nil
}

func (h *AuthHelper) HashPassword(password string) (string, error) {
	h.log.Debug().Msg("Hashing password")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *AuthHelper) CheckPasswordHash(password, hash string) bool {
	h.log.Debug().Msg("Checking password hash")
	if password == "" || hash == "" {
		h.log.Error().Msg("Password or hash is empty")
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
