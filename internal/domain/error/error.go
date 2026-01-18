package app_err

import "github.com/lubosgarancovsky/go-kit"

var (
	ErrInsufficientPermission = go_kit.ErrForbidden.WithMessage("you don't have sufficient permission for this action")
)
