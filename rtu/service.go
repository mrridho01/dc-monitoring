package rtu

import (
	"dc-monitor/analogs"
	"dc-monitor/ghs"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Establish connection to RTU, retrieve data modbus, store to database
func RetrieveModbus(address string, slave byte, pollingtime time.Duration) {
	// Establish RTU connection
	tcpClientHandler, client := createClientHandler(address, slave)
	e := tcpClientHandler.Connect()
	if e != nil {
		log.Panic("Cannot establish connection to RTU: ", e)
		return
	}
	defer tcpClientHandler.Close()

	//Establish Database Connection
	db, err := gorm.Open(sqlite.Open("entity/test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database: ", err)
		return
	}

	//Migrate schema
	db.AutoMigrate(&ghs.GH{})
	db.AutoMigrate(&analogs.Analog{})

	// Query Modbus Slave, store to database
	db.Create(&ghs.GH{Name: "Cemara"})
	for range time.Tick(pollingtime * time.Second) {
		current1, _ := client.ReadInputRegisters(1, 1)
		analog := &analogs.InputRegister{
			AcInput:   float32(current1[1]),
			Dc110:     float32(current1[1]),
			Dc48:      float32(current1[1]),
			CurrentAc: float32(current1[1]),
			CurrentDc: float32(current1[1]),
		}
		db.Create(&analogs.Analog{
			GHID:          1,
			InputRegister: *analog,
		})
	}
}
