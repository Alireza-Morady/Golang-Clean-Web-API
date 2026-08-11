package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHnadler() *HealthHandler{
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,"Working!!")
	return
}
func (h *HealthHandler) HealthById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK,fmt.Sprintf("Post by id : %s is Working!!",id))
	return
}
