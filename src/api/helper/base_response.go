package helper

import "github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/validation"

type BaseHttpResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"resultCode"`
	ValidationError *[]validation.ValidationError `json:"validationError"`
	Error            any                           `json:"error"`
}

func GenerateBaseResponse(result any, success bool,resultCode int) *BaseHttpResponse{
	return &BaseHttpResponse{
		Result: result,
		Success: success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponsiveWithError(result any, success bool,resultCode int,err error)*BaseHttpResponse{
	return &BaseHttpResponse{
		Result: result,
		Success: success,
		ResultCode: resultCode,
		Error: err,
	}
}

func GenerateBaseResponsiveWithValidationError(result any, success bool,resultCode int,err error,ve validation.ValidationError)*BaseHttpResponse{
	return &BaseHttpResponse{
		Result: result,
		Success: success,
		ResultCode: resultCode,
		Error: err.Error(),
		ValidationError: validation.GetValidationError(err),
	}
}