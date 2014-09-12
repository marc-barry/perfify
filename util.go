package main

import (
	"net"
	"os"
)

func closeTCPListener(listener *net.TCPListener) {
	err := listener.Close()

	if err != nil {
		println("Close failed: ", err.Error())
		os.Exit(1)
	}
}

func closeTCPConn(conn *net.TCPConn) {
	err := conn.Close()

	if err != nil {
		println("Close failed: ", err.Error())
		os.Exit(1)
	}
}

func closeUDPConn(conn *net.UDPConn) {
	err := conn.Close()

	if err != nil {
		println("Close failed: ", err.Error())
		os.Exit(1)
	}
}
