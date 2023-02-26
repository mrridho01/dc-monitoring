package rtu

import (
	"log"
	"os"
	"time"

	"github.com/goburrow/modbus"
)

type Client struct {
	Address string
	SlaveId byte
	Timeout time.Duration
	Logger  *log.Logger
}

func createClientHandler(address string, id byte) (*modbus.TCPClientHandler, modbus.Client) {
	clientConfig := Client{
		Address: address,
		SlaveId: id,
		Timeout: 10 * time.Second,
		Logger:  log.New(os.Stdout, "", log.Ldate),
	}

	handler := modbus.NewTCPClientHandler(clientConfig.Address)
	handler.Timeout = clientConfig.Timeout
	handler.SlaveId = clientConfig.SlaveId
	handler.Timeout = clientConfig.Timeout
	handler.Logger = clientConfig.Logger

	client := modbus.NewClient(handler)
	return handler, client
}
