package common

import "github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"

var (
	lowerCharSet   = "abcdefghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%^*"
	numberSet      = "0123456789"
	allCharSet	   = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func CheckPassword(password string) bool {
	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLength{
		return false
	}

	return true
}
