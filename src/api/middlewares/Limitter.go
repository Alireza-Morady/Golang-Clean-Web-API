package middlewares

import (
	"net/http"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest()gin.HandlerFunc{
	limiter := tollbooth.NewLimiter(1,nil)
	return func(c *gin.Context){
		err:= tollbooth.LimitByRequest(limiter,c.Writer,c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,helper.GenerateBaseResponsiveWithError(nil,false,http.StatusTooManyRequests,err))
			return 
		}
		c.Next()
	}
}