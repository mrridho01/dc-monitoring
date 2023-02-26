package main

import (
	"dc-monitor/web"
	"log"
)

func main() {
	// log.Println("Starting telegram service...")
	// go startTelegram()

	// log.Println("Starting modbus master service...")
	// rtu.RetrieveModbus("192.168.100.102:502", 1, 5)
	log.Println("Starting web service")
	web.StartServer()
}
