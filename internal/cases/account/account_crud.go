package account

import (
	"condomilux/condo-admin/internal/helpers"
	"condomilux/condo-admin/internal/models"
	"context"

	"github.com/carlware/go-common/errors"
)

// Create in a comment
func Create(ctx context.Context, opts *Opts, req *CreateRequest) (*models.Account, error) {
	oa := "cases.account.Create"

	if req == nil {
		return nil, helpers.NewNilErr(oa)
	}

	account := &models.Account{}
	if err := helpers.Copy(account, req); err != nil {
		return nil, helpers.NewInternalErr(oa, err)
	}

	account.Initialize()

	if err := account.Validate(); err != nil {
		return nil, helpers.NewArgsErr(oa, err)
	}

	p, err := opts.Repository.Add(context.Background(), account)
	if err != nil {
		return nil, errors.New(errors.Internal, oa, "", err)
	}

	return p, nil
}

// Retrieve in a comment
func Retrieve(ctx context.Context, opts *Opts, id string) (*models.Account, error) {
	oa := "cases.account.Retrieve"

	if id == "" {
		return nil, helpers.NewEmptyIDErr(oa)
	}

	p, err := opts.Repository.Get(context.Background(), id)
	if err != nil {
		return nil, errors.New(errors.Internal, oa, "", err)
	}

	return p, nil
}

// Update in a comment
func Update(ctx context.Context, opts *Opts, id string, req *UpdateRequest) (*models.Account, error) {
	oa := "cases.account.Update"

	if req == nil {
		return nil, helpers.NewNilErr(oa)
	}

	account, err := opts.Repository.Get(context.Background(), id)
	if err != nil {
		return nil, errors.New(errors.Internal, oa, "", err)
	}

	if err := helpers.UpdateStruct(account, &models.Account{}, req); err != nil {
		return nil, err
	}

	if err := account.Validate(); err != nil {
		return nil, helpers.NewArgsErr(oa, err)
	}

	p, err := opts.Repository.Update(context.Background(), account)
	if err != nil {
		return nil, errors.New(errors.Internal, oa, "", err)
	}
	return p, nil
}

// Delete in a comment
func Delete(ctx context.Context, opts *Opts, id string) (*models.Account, error) {
	oa := "cases.account.Delete"

	if id == "" {
		return nil, helpers.NewEmptyIDErr(oa)
	}

	account := &models.Account{ID: id}

	p, err := opts.Repository.Remove(context.Background(), account)
	if err != nil {
		return nil, errors.New(errors.Internal, oa, "", err)
	}

	return p, nil
}

// List in a comment
func List(cxt context.Context, opts *Opts) ([]*models.Account, error) {
	oa := "cases.account.List"

	ps, err := opts.Repository.List(context.Background())
	if err != nil {
		return nil, errors.New(errors.Internal, oa, "", err)
	}

	return ps, nil
}
