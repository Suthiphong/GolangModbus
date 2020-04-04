package main

import (
	"fmt"
	"time"

	"github.com/goburrow/modbus"
)

func main() {
	fmt.Println("OK RUNNING.")
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 2
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	for {
		results, err := client.ReadDiscreteInputs(0, 10)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("result time : ", time.Now())
		fmt.Println(results)
		time.Sleep(1 * time.Second)
	}
}
