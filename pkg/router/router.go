package router

import (
	"election_algorithm/pkg/router/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/health-update", handler.HealthUpdate)

	return router
}
