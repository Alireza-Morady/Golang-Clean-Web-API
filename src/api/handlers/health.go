package handlers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHnadler() *HealthHandler{
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,"yeah Working!!")
}

