package helpers

import (
	"github.com/kobeld/goutils"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pwd string) (encryptedPwd string, err error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(pwd), 0)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}
	encryptedPwd = string(hp)
	return
}

func IsPwdMatch(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}
