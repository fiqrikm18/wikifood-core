package user

import (
	"WikfoodCore/internal/infrastuctures/persistence"
	"context"
)

type UserRepository struct {
	dbConf *persistence.DBConnection
	ctx    context.Context
}

func New(ctx context.Context, db *persistence.DBConnection) *UserRepository {
	return &UserRepository{
		dbConf: db,
		ctx:    ctx,
	}
}
