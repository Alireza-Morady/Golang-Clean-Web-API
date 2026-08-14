package validation

import (
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/common"
	"github.com/go-playground/validator/v10"
)



func IranMobileNumberValidator(fld validator.FieldLevel)bool{
	value,ok:= fld.Field().Interface().(string)
		if !ok{
			return false
		}
		return common.IranMobileNumberValidate(value)
}		