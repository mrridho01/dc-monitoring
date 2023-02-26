package handler

import (
	"dc-monitor/ghs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// handler dashboard html template
func DashboardHandler(c *gin.Context) {
	//Query semua GH dari DB, preload analogs terakhir
	db, err := gorm.Open(sqlite.Open("entity/test.db"), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusNotFound, "404 Status Not Found")
	}
	var ghs []ghs.GH
	db.Preload("Analogs", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC").Limit(1)
	}).Find(&ghs)

	// Return HTML template
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"gh": ghs,
	})

}
