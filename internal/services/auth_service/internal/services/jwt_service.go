package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
)

type TokenManager interface {
	NewAccessToken(userId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	secretKey string
	logger    logger.ILogger
}

func NewManager(tknCfg *config.TokenManagerConfig, logger logger.ILogger) *Manager {
	return &Manager{secretKey: tknCfg.SecretKey, logger: logger}
}

func (m *Manager) NewAccessToken(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.secretKey))
}

func (m *Manager) NewRefreshToken() (string, error) {
	buffer := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(buffer)
	if err != nil {
		m.logger.Error("error during creation of refresh token: " + err.Error())
		return "", err
	}

	return fmt.Sprintf("%x", buffer), nil
}
