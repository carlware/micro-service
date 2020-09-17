package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"condomilux/condo-admin/cli/dispatchers/graphql/graph/generated"
	"condomilux/condo-admin/cli/dispatchers/graphql/graph/model"
	"condomilux/condo-admin/internal/cases/account"
	"condomilux/condo-admin/internal/models"
	"context"

	e "github.com/carlware/go-common/errors"
	"github.com/carlware/go-common/log"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)


func (r *mutationResolver) CreateAccount(ctx context.Context, input account.CreateRequest) (*models.Account, error) {
	opts := &account.Opts{r.Ctrl.Repositories.Account}

	res, err := account.Create(ctx, opts, &input)
	if err != nil {
		log.For(ctx).Error("mutation error", zap.Error(err))
		return nil, xerrors.New(e.ErrorMessage(err))
	}
	return res, nil
}


// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
