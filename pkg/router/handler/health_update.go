package handler

import (
	"election_algorithm/pkg/model"
	"election_algorithm/pkg/service"
	"github.com/gin-gonic/gin"
)

func HealthUpdate(c *gin.Context) {
	var req model.HealthUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}

	go service.HealthUpdate(&req)

	c.JSON(204, nil)
}
