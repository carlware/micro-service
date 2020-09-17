package interfaces

import (
	"condomilux/condo-admin/internal/models"
	"context"

	_ "github.com/golang/mock/mockgen/model" // I justify
)

//go:generate mockgen -destination mocks/account.repository.gen.go -package mocks condomilux/condo-admin/internal/interfaces Account
type Account interface {
	Add(ctx context.Context, account *models.Account) (*models.Account, error)
	Remove(ctx context.Context, account *models.Account) (*models.Account, error)
	Get(ctx context.Context, id string) (*models.Account, error)
	Update(ctx context.Context, account *models.Account) (*models.Account, error)
	List(ctx context.Context) ([]*models.Account, error)
}
