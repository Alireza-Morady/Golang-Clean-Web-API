package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func TestMiddleWare()gin.HandlerFunc{
	return func(ctx *gin.Context){
		apikey := ctx.GetHeader("x-api-key")
		if apikey == "1"{
			ctx.Next()
			return 
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
			"result": "API key is required",
		})
		
	}

}