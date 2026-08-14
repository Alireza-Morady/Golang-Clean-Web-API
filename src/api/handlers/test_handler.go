package handlers

import (
	"net/http"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/helper"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/validation"
	"github.com/gin-gonic/gin"
)

type header struct {
	UserId string
}

type PersonData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}
func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}
func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}
func (h *TestHandler) UserByname(c *gin.Context) {
	username := c.Params.ByName("username")
	c.JSON(http.StatusOK, gin.H{
		"result":   "UserByname",
		"username": username,
	})
}
func (h *TestHandler) Accounts(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     id,
	})
}
func (h *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
	})
}
func (h *TestHandler) HeaderBinder1(c *gin.Context) {

	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"UserId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"UserId": header,
	})

}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"UserId": id,
		"name":   name,
	})

}
func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result":  "QueryBinder2",
		"UsersId": ids,
		"name":    name,
	})
}
func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result":  "UriBinder",
		"UsersId": id,
		"name":    name,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := PersonData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,helper.GenerateBaseResponsiveWithValidationError("bad request",false,http.StatusBadRequest,err,validation.ValidationError{}))
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": "BodyBinder",
	// 	"person": p,
	// })
	c.JSON(http.StatusOK,helper.GenerateBaseResponse(p,true,200))

}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := PersonData{}
	c.Bind(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"person": p,
	})
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "file")
	c.JSON(http.StatusOK, gin.H{
		"result": "FileBinder",
		"person": file.Filename,
	})
}
