package entity

import "errors"

var (
	ErrUserAlredyExists = errors.New("user with such email or phone already exists")
	
)