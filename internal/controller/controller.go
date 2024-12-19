package controller

import (
	"docker-example/internal/controller/dto"
	"docker-example/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	router *gin.Engine
}

func Init(router *gin.Engine) *Controller {
	controller := &Controller{
		router,
	}
	controller.router.GET("/", controller.getRecordCount)
	controller.router.GET("/records", controller.getRecords)
	controller.router.POST("/records", controller.addRecord)
	controller.router.DELETE("/records", controller.removeRecords)

	return controller
}

func (controler *Controller) getRecordCount(ctx *gin.Context) {
	recordCount, err := service.GetRecordCount()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"recordCount": recordCount})
}
func (controler *Controller) getRecords(ctx *gin.Context) {
	records, err := service.GetRecords()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"records": records})
}
func (controler *Controller) addRecord(ctx *gin.Context) {
	dto := new(dto.AddRecord)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	record, err := service.AddRecord(dto.Text)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"newRecord": record})
}
func (controler *Controller) removeRecords(ctx *gin.Context) {
	err := service.RemoveRecords()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusNoContent)
}
