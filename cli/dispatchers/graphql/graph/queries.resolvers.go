package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"arquil/accounts/cli/dispatchers/graphql/graph/generated"
	"arquil/accounts/internal/cases/account"
	"arquil/accounts/internal/models"
	"context"
	"fmt"

	e "github.com/carlware/go-common/errors"
	"github.com/carlware/go-common/log"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

func (r *queryResolver) Account(ctx context.Context, id string) (*models.Account, error) {
	opts := &account.Opts{r.Ctrl.Repositories.Account}

	res, err := account.Retrieve(ctx, opts, id)
	if err != nil {
		log.For(ctx).Error("query error", zap.Error(err))
		return nil, xerrors.New(e.ErrorMessage(err))
	}
	return res, nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*models.Account, error) {
	opts := &account.Opts{r.Ctrl.Repositories.Account}

	res, err := account.List(ctx, opts)
	if err != nil {
		log.For(ctx).Error("query error", zap.Error(err))
		return nil, xerrors.New(e.ErrorMessage(err))
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
