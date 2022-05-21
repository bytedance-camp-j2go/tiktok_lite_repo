package base

import (
	"errors"
)

var (
	ErrPathNotFound = errors.New("path not found")
	ErrNotFile      = errors.New("not file")
	ErrNotImplement = errors.New("not implement")
	ErrNotSupport   = errors.New("not support")
	ErrNotFolder    = errors.New("not a folder")
	ErrEmptyFile    = errors.New("empty file")
	ErrRelativePath = errors.New("access using relative path is not allowed")
	ErrEmptyToken   = errors.New("empty token")
)
