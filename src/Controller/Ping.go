package controller

import (
	model "example/socket/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct{}

func (ping *Ping) Init(group *gin.RouterGroup, db map[string]model.User) {
	v1 := group.Group("/v1")

	pingGroup := v1.Group("/ping")
	pingGroup.GET("/", ping.getPing)
}

func (ping *Ping) getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
