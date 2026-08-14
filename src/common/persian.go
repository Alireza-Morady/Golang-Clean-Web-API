package common

import (
	"log"
	"regexp"
)

const iraniamMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranMobileNumberValidate(value string)bool{
	result,err :=regexp.MatchString(iraniamMobileNumberPattern,value)
	if err != nil{
		log.Print(err.Error())
	}
	return result
}