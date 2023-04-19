package mp

import "errors"

var (
	ErrInvalidKey   = errors.New("key should not be empty")
	ErrInvalidValue = errors.New("value should be map[string]interface{}")
	ErrInvalidParam = errors.New("invalid parameter")
)
