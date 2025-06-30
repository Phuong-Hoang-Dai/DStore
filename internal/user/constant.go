package user

import "errors"

const MaxLimit = 50
const UserTableName = "User"

var (
	ErrUserNameOrPasswordIncorrect    = errors.New("username or password is incorrect")
	ErrAuthorizationHeaderMissing     = errors.New("authorization header is missing")
	ErrAuthorizationHeaderWrongFormat = errors.New("authorization header is wrong format")
	ErrMissingRole                    = errors.New("missing role")
	ErrNotAllowToAccess               = errors.New("not allow to access")
)
