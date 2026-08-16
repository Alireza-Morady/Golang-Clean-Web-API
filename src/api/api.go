package api

import (
	// "net/http"

	"fmt"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/middlewares"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/routers"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/validation"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	
	r := gin.New()

	RegisterValidators()

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest()/*, middlewares.TestMiddleWare()*/)

	RegisterRoutes(r)
	RegisterSwagger(r,cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}


func RegisterValidators(){
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}
func RegisterRoutes(r *gin.Engine){
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.TestRouters(test_router)
		routers.Health(health)
	}
}
func RegisterSwagger(r *gin.Engine,cfg *config.Config){
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s",cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
}