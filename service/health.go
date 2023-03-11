package service

import (
	"context"

	"github.com/kwtryo/user-management-kawata-api/store"
)

type HealthRepository interface {
	Ping(ctx context.Context, db store.DBConnection) error
}

type HealthService struct {
	DB   store.DBConnection
	Repo HealthRepository
}

func (hs *HealthService) HealthCheck(ctx context.Context) error {
	return hs.Repo.Ping(ctx, hs.DB)
}
