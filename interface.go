package main

import (
	"fmt"
	"net"
)

func printInterfaces() {
	if interfaces, err := net.Interfaces(); err != nil {
		println(err.Error())
	} else {
		for _, interf := range interfaces {
			fmt.Printf("Index: %d, MTU: %d, Name: %s, Addr: %s, Flags: %s \n", interf.Index, interf.MTU, interf.Name, interf.HardwareAddr.String(), interf.Flags.String())
		}
	}
}
