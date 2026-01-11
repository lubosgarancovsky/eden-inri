package app_err

import "github.com/lubosgarancovsky/go-kit"

var (
	ErrNotAMember              = go_kit.ErrForbidden.WithMessage("User is not a member of the project")
	ErrInsufficientProjectRole = go_kit.ErrForbidden.WithMessage("User does not have sufficient role in the project")
)
