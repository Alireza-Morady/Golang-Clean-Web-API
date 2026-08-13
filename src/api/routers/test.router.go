package routers

import (
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/handlers"
	"github.com/gin-gonic/gin"
)


func TestRouters(r *gin.RouterGroup){
	h := handlers.NewTestHandler()
	r.GET("/",h.Test)
	r.GET("/users",h.Users)
	r.GET("/user/:id",h.UserById)
	r.GET("/user/get-userby-username/:username",h.UserByname)
	r.GET("/user/:id/accounts",h.Accounts)
	r.POST("/user/:id/accounts",h.AddUser)

	r.POST("/HeaderBinder1",h.HeaderBinder1)
	r.POST("/HeaderBinder2",h.HeaderBinder2)

	r.POST("/query1",h.QueryBinder1)
	r.POST("/query2",h.QueryBinder2)
	r.POST("/Uri/:id/:name",h.UriBinder)
	r.POST("/binder/body",h.BodyBinder)
	r.POST("/binder/form",h.FormBinder)
	r.POST("/binder/file",h.FileBinder)
}