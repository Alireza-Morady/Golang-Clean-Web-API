package api

import (
	// "net/http"

	"fmt"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/middlewares"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/routers"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/validation"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(), middlewares.TestMiddleWare())
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.TestRouters(test_router)
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
