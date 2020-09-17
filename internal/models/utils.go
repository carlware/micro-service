package models

import "github.com/carlware/go-common/errors"

type Status struct {
	Code    errors.ErrorCode `json:"code"`
	Message string           `json:"message"`
}
