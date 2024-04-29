package controller

import (
	"github.com/gin-gonic/gin"
	"invoice/process/src/application/service"
	"net/http"
)

type BatchController struct {
	service *service.BatchService
}

func NewBatchController(service *service.BatchService) *BatchController {
	return &BatchController{
		service: service,
	}
}

func (b *BatchController) InitRoutes() {
	app := gin.Default()
	api := app.Group("/api/process")

	api.POST("/", b.migrate)
	app.Run(":3000")
}

type PathRequest struct {
	Path string `json:"path"`
}

func (b *BatchController) migrate(ctx *gin.Context) {
	var pathRequest PathRequest

	err := ctx.ShouldBindJSON(&pathRequest)
	if err != nil {
		return
	}
	b.service.Migrate(pathRequest.Path)
	ctx.JSON(http.StatusOK, gin.H{"path": pathRequest})
}
