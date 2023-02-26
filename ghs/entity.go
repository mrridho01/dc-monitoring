package ghs

import (
	"dc-monitor/analogs"

	"gorm.io/gorm"
)

type GH struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Analogs []analogs.Analog
}
