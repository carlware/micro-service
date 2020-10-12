package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"carlware/accounts/cli/dispatchers/graphql/graph/generated"
	"carlware/accounts/internal/cases/account"
	"carlware/accounts/internal/models"
	"context"
	"fmt"

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

func (r *mutationResolver) UpdateAccount(ctx context.Context, id string, input account.UpdateRequest) (*models.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, id string) (*models.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
