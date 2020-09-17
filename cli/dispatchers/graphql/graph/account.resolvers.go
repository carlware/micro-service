package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"arquil/accounts/cli/dispatchers/graphql/graph/generated"
	"arquil/accounts/internal/models"
	"context"
)

func (r *accountResolver) ID(ctx context.Context, obj *models.Account) (string, error) {
	return string(obj.ID), nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
