package handler

import (
	"dc-monitor/analogs"
	"dc-monitor/ghs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// get all data for query to "/"
func APIDashboardHandler(c *gin.Context) {
	// Query jumlah GH dari DB
	db, err := gorm.Open(sqlite.Open("entity/test.db"), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusNotFound, "404 Status Not Found")
	}
	var ghs []ghs.GH
	db.Find(&ghs)

	//Query nilai Analog dari DB
	var analogs []analogs.Analog
	for _, gh := range ghs {
		db.Where("gh_id = ?", gh.ID).Last(&analogs)
	}

	c.JSON(http.StatusOK, analogs)

}

// get gh last record
func APIGHHandler(c *gin.Context) {
	// connect to database
	db, err := gorm.Open(sqlite.Open("entity/test.db"), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusNotFound, "404 Status Not Found")
	}

	// query last record of GH
	var analogs []analogs.Analog
	id := c.Param("id")
	analogData := db.Where("gh_id = ?", id).Last(&analogs)
	if analogData.Error != nil {
		c.JSON(http.StatusInternalServerError, analogs)
	} else {
		c.JSON(http.StatusOK, analogs)
	}
}
