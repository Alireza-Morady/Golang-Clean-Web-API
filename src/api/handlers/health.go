package handlers

import (
	// "fmt"
	"net/http"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHnadler() *HealthHandler{
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,helper.GenerateBaseResponse("working",true,0))
}

