package cache

import "errors"

var (
	ErrNotFound = errors.New("Not found")
	ErrExists   = errors.New("Error exists")
)
