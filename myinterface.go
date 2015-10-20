package main

import (
	"fmt"
)

//任何struct都实现了空接口
type empty interface {
}

type USB interface {
	Name() string
	Connect()
}

//实现了接口的所以方法，便是实现了这个接口
type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("connected", pc.name)
}

func main() {
	var usb USB
	pc := PhoneConnector{"PConnector"}
	usb = USB(pc)
	usb.Connect()
}
