package domain

import "errors"

var (
	ErrResourceNotFound  = errors.New("resource not found")
	ErrUserAlreadyExists = errors.New("user with this credential already exists")
	ErrWrongCredentials  = errors.New("credentials are not valid")
	ErrEmptyID           = errors.New("id param is empty")
)
