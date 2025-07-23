package cep

import "errors"

type ErrorCep error

var (
	ErrInvalidCep  ErrorCep = errors.New("invalid zopcode")
	ErrCepNotFound ErrorCep = errors.New("can not find zipcode")
)
