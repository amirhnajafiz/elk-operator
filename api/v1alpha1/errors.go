package v1alpha1

import "errors"

var (
	ErrUsernameExists = errors.New("cannot register your user, username is duplicated")
	ErrUserNotFound   = errors.New("user with the requested username not found")
	ErrReplicaNumber  = errors.New("cannot set even values for replica number, use odd numbers")
)
