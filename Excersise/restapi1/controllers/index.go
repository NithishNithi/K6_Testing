package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Junk1(c *gin.Context){
	c.String(http.StatusOK,"Pong")
	
}
