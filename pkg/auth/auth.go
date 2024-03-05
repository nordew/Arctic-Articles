package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nordew/ArcticArticles/pkg/logging"
	"strings"
	"time"
)

type jwtAuthenticator struct {
	signKey string
	logger  logging.Logger
}

func NewAuth(signKey string, logger logging.Logger) Authenticator {
	return &jwtAuthenticator{
		signKey: signKey,
		logger:  logger,
	}
}

type TokenClaims struct {
	UserId string `json:"sub"`
	Role   int    `json:"role"`
	jwt.RegisteredClaims
}

func (s *jwtAuthenticator) GenerateTokens(options *GenerateTokenClaimsOptions) (string, string, error) {
	const op = "jwt.GenerateTokens"

	mySigningKey := []byte(s.signKey)

	claims := TokenClaims{
		UserId: options.UserId,
		Role:   options.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "arctic-api",
			Subject:   "client",
			ID:        uuid.NewString(),
			Audience:  []string{"arctic-articles"},
		},
	}

	refreshToken, err := s.GenerateRefreshToken(options.UserId, options.Role)
	if err != nil {
		s.logger.Error("failed to generate refresh token", err.Error(), op)
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(mySigningKey)
	if err != nil {
		s.logger.Error("failed to sign token", err.Error(), op)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *jwtAuthenticator) GenerateRefreshToken(id string, role int) (string, error) {
	const op = "jwt.GenerateRefreshToken"

	mySigningKey := []byte(s.signKey)

	claims := TokenClaims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "arctic-api",
			Subject:   "client",
			ID:        uuid.NewString(),
			Audience:  []string{"arctic-articles"},
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedRefreshToken, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		s.logger.Error("failed to sign refresh token", err.Error(), op)
		return "", err
	}

	return signedRefreshToken, nil
}

func (s *jwtAuthenticator) ParseToken(accessToken string) (*ParseTokenClaimsOutput, error) {
	const op = "jwt.ParseToken"

	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signKey), nil
	})
	if err != nil {
		s.logger.Error("failed to parse token", err.Error(), op)
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}

	if !token.Valid {
		s.logger.Error("token is not valid", op)
		return nil, fmt.Errorf("token is not valid")
	}

	claims := token.Claims.(jwt.MapClaims)

	role := claims["role"]
	if role == nil {
		s.logger.Error("token is not valid: missing role")
		return nil, fmt.Errorf("token is not valid")
	}
	sub := claims["sub"]
	if sub == nil {
		s.logger.Error("token is not valid: missing subject")
		return nil, fmt.Errorf("token is not valid")
	}

	roleVal, ok := role.(int)
	if !ok {
		s.logger.Error("token role is invalid not int")
		return nil, fmt.Errorf("token role is invalid not int")
	}

	return &ParseTokenClaimsOutput{Sub: fmt.Sprint(sub), Role: roleVal}, nil
}
