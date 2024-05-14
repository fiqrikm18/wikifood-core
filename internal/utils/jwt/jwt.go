package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	secret = ""
)

type JwtClaimsPayload struct {
	AccesTokenExp   time.Time
	RefreshTokenExp time.Time
	AccessTokenId   uuid.UUID
	RefreshTokenId  uuid.UUID
	Issuer          string
	UserId          uint
}

type JWTToken struct{}

func NewJWTToken() *JWTToken {
	return &JWTToken{}
}

// GenerateToken generating jwt token base on claims as payload
func (j *JWTToken) GenerateToken(jwtClaims JwtClaimsPayload) (string, error) {
	claims := jwt.MapClaims{
		"access_token_exp":  jwtClaims.AccesTokenExp,
		"refresh_token_exp": jwtClaims.RefreshTokenExp,
		"access_token_id":   jwtClaims.AccessTokenId,
		"refresh_token_id":  jwtClaims.RefreshTokenId,
		"issuer":            jwtClaims.Issuer,
		"user_id":           jwtClaims.UserId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validating token return true if token is qualified to use and false if not qualified to use
func (j *JWTToken) ValidateToken(tokenString string, tokenType TokenType) (bool, error) {
	tokenClaims, err := j.Parse(tokenString)
	if err != nil {
		return false, err
	}

	if tokenType == ACCESS_TOKEN {
		if tokenClaims.AccesTokenExp.After(time.Now()) {
			return false, errors.New("token expired")
		}
	} else {
		if tokenClaims.RefreshTokenExp.After(time.Now()) {
			return false, errors.New("token expired")
		}
	}

	return true, nil
}

// Parse parse token string and return token claims
func (j *JWTToken) Parse(tokenString string) (*JwtClaimsPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid string token")
	}

	return &JwtClaimsPayload{
		claims["access_token_exp"].(time.Time),
		claims["refresh_token_exp"].(time.Time),
		claims["access_token_id"].(uuid.UUID),
		claims["refresh_token_id"].(uuid.UUID),
		claims["issuer"].(string),
		claims["user_id"].(uint),
	}, nil
}
