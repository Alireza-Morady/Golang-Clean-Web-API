package validation

import (
	"log"
	"regexp"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/common"
	"github.com/go-playground/validator/v10"
)

// ^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$
func IranMobileNumberValidator(fld validator.FieldLevel)bool{
value,ok:= fld.Field().Interface().(string)
	if !ok{
		return false
	}
	result,err :=regexp.MatchString(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`,value)
	if err != nil{
		log.Print(err.Error())
	}
	return result
}

func PasswordValidator(fld validator.FieldLevel)bool{
	value,ok:= fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}
	return common.CheckPassword(value)
}