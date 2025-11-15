package refresh_token

import "errors"

var (
	ErrorThisUserRefreshTokenNotFound = errors.New("this user's refresh token not found")
	ErrorPasswordNotMatch             = errors.New("password does not match")
)
