package global

import (
	"fmt"

	"gopkg.in/mgo.v2"

	"errors"
)

var (
	ErrInvalidId         error = errors.New("Invalid ObjectId")
	ErrInvalidParam      error = errors.New("Invalid Parameter")
	ErrInvalidTime       error = errors.New("Invalid Time Selection")
	ErrInvalidTimeFormat error = errors.New("Invalid Time Format")
	ErrInvalidTimestamp  error = errors.New("Invalid Timestamp")
	ErrInvalidUpload     error = errors.New("Invalid Upload")
	ErrInvalidLimit      error = errors.New("Should provide a valid number")
	ErrPermissionDenied  error = errors.New("Permission Denied")
	ErrPwdNotMatch       error = errors.New("Password doesn't match")
	ErrAccountTaken      error = errors.New("The mobile has been taken")
	ErrLoginFailed       error = errors.New("Account and Password don't match")
	ErrUserNotFound      error = errors.New("User not found")
	ErrNotFound          error = errors.New("Not found")
	ErrShouldBeBlank     error = errors.New("Should be blank")
	ErrShouldNotBebBlank error = errors.New("Should not be blank")

	ErrDeviceTaken error = errors.New("The device has been registered")

	// For Share Request
	ErrInvalidToken error = errors.New("Invalid Token")
	ErrTokenExpired error = errors.New("Token Expired")

	ErrInvalidEnv      error = errors.New("Runtime Env (dev/prod/test/ci) is not provided")
	ErrServiceNotFound error = errors.New("Service not found")

	// For Category
	ErrAliasIsTaken error = errors.New("The Alias is being taken")

	ErrMgoNotFound error = mgo.ErrNotFound

	ErrCannotFollowYourSelf error = errors.New("Can't follow the same user")
)

func BuildNotFoundError(str string) error {
	return errors.New(fmt.Sprintf("%s is not found", str))
}
