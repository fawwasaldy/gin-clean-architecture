package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type (
	JWTService interface {
		GenerateAccessToken(userID string, role string) string
		GenerateRefreshToken() (string, time.Time)
		ValidateToken(token string) (*jwt.Token, error)
		GetUserIDByToken(token string) (string, error)
	}

	jwtCustomClaim struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
		jwt.RegisteredClaims
	}

	jwtService struct {
		secretKey         string
		issuer            string
		accessExpiration  time.Duration
		refreshExpiration time.Duration
	}
)

func NewJWTService() JWTService {
	return &jwtService{
		secretKey:         getSecretKey(),
		issuer:            getIssuer(),
		accessExpiration:  getAccessExpiration(),
		refreshExpiration: getRefreshExpiration(),
	}
}

func (j *jwtService) GenerateAccessToken(userID string, role string) string {
	claims := jwtCustomClaim{
		userID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessExpiration)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}

	return tokenString
}

func (j *jwtService) GenerateRefreshToken() (string, time.Time) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
		return "", time.Time{}
	}

	refreshToken := base64.StdEncoding.EncodeToString(b)
	expiresAt := time.Now().Add(j.refreshExpiration)

	return refreshToken, expiresAt
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}

func (j *jwtService) GetUserIDByToken(token string) (string, error) {
	parsedToken, err := j.ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID := fmt.Sprintf("%v", claims["user_id"])
	return userID, nil
}

func (j *jwtService) parseToken(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "kpl-base-secret"
	}
	return secretKey
}

func getIssuer() string {
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "kpl-base"
	}
	return issuer
}

func getAccessExpiration() time.Duration {
	expiration := os.Getenv("JWT_ACCESS_EXPIRATION")
	if expiration == "" {
		expiration = "15m"
	}
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		duration = 15 * time.Minute
	}
	return duration
}

func getRefreshExpiration() time.Duration {
	expiration := os.Getenv("JWT_REFRESH_EXPIRATION")
	if expiration == "" {
		expiration = "7d"
	}
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		duration = 7 * 24 * time.Hour
	}
	return duration
}
