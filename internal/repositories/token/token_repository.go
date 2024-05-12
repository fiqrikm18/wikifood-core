package token

import (
	"WikfoodCore/internal/infrastuctures/persistence"
	"context"
)

type TokenRepository struct {
	dbConf *persistence.DBConnection
	ctx    context.Context
}

func New(ctx context.Context, db *persistence.DBConnection) *TokenRepository {
	return &TokenRepository{
		dbConf: db,
		ctx:    ctx,
	}
}
