package token

import "WikfoodCore/internal/repositories/token"

type TokenService struct {
	tokenRepository token.ITokenRepository
}

func New(repository token.ITokenRepository) *TokenService {
	return &TokenService{
		tokenRepository: repository,
	}
}
