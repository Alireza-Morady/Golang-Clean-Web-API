package common

import (
	// "log"
	"regexp"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/pkg/logging"
)
var logger = logging.NewLogger(config.GetConfig())
const iraniamMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranMobileNumberValidate(value string)bool{
	result,err :=regexp.MatchString(iraniamMobileNumberPattern,value)
	if err != nil{
		logger.Error(logging.Validation,logging.MobileValidation,err.Error(),nil)
	}
	return result
}