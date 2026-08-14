package routers

import (
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/handlers"
	"github.com/gin-gonic/gin"
)
func Health(r *gin.RouterGroup){
	h := handlers.NewHealthHnadler()

	r.GET("/",h.Health)
}