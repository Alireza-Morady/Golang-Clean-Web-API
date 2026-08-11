package routers

import (
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/handlers"
	"github.com/gin-gonic/gin"
)


func Health(r *gin.RouterGroup){

	handler := handlers.NewHealthHnadler()

	r.GET("/",handler.Health)
	r.POST("/:id",handler.HealthById)
}