package controller

import (
	"docker/internal/service"
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
	controller.router.DELETE("/records", controller.deleteRecords)

	return controller
}

func (controler *Controller) getRecordCount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"recordCount": service.GetRecordCount()})
}
func (controler *Controller) getRecords(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"records": service.GetRecords()})
}
func (controler *Controller) addRecord(ctx *gin.Context) {
	req := new(struct {
		Text string `json:"text" binding:"required"`
	})
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"newRecord": service.AddRecord(req.Text)})
}
func (controler *Controller) deleteRecords(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"recordCount": service.DeleteRecords()})
}
