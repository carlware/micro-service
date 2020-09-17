package helpers

import (
	"github.com/carlware/go-common/errors"
)

func NewInternalErr(op string, err error) error {
	return errors.New(errors.Internal, op, err.Error(), nil)
}

func NewNilErr(op string) error {
	return errors.New(errors.InvalidArgument, op, "nil request error", nil)
}

func NewEmptyIDErr(op string) error {
	return errors.New(errors.InvalidArgument, op, "id cannot be blank", nil)
}

func NewNilRepositoryErr(op string) error {
	return errors.New(errors.Internal, op, "nil repository", nil)
}

func NewArgsErr(op string, err error) error {
	e := errors.As(err)
	return errors.New(errors.InvalidArgument, op, e.Error(), err)
}

func NewDatabaseError(op string, err error) error {
	switch err.Error() {
	case "pg: no rows in result set":
		return errors.New(errors.NotFound, op, "A resource with this ID does not exists", err)
	default:
		return errors.New(errors.Internal, op, "", err)
	}
}
