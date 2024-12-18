package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Record struct {
	RecordId *uuid.UUID `json:"record_id"`
	Text     string     `json:"text"`
}

var record = []*Record{}

func main() {

	router := gin.Default()
	router.GET("/", getRecordCount)
	router.GET("/records", getRecords)
	router.POST("/records", addRecord)
	router.DELETE("/records", deleteRecords)
	router.Run(":8000")

}
func getRecordCount(c *gin.Context) {
	log.Printf("getRecordCount\n")
	c.JSON(http.StatusOK, gin.H{"recordCount": len(record)})
}
func getRecords(c *gin.Context) {
	log.Printf("getRecords\n")
	c.JSON(http.StatusOK, gin.H{"records": record})
}
func addRecord(c *gin.Context) {
	req := new(struct {
		Text string `json:"text"`
	})
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": err})
	}
	recordId := uuid.New()
	newRecord := &Record{
		RecordId: &recordId,
		Text:     req.Text,
	}
	record = append(record, newRecord)
	log.Printf("addRecord\n")
	c.JSON(http.StatusOK, gin.H{"newRecord": newRecord})
}
func deleteRecords(c *gin.Context) {
	record = []*Record{}
	log.Printf("deleteRecords\n")
	c.JSON(http.StatusOK, gin.H{"recordCount": 0})
}
