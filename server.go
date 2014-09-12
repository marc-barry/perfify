package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"net"
	"os"
)

func tcpServerCommand(c *cli.Context) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", DEFAULT_SERVER_ADDR)
	if err != nil {
		println("Resolving address failed: ", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		println("Dial failed: ", err.Error())
		os.Exit(1)
	}

	defer closeTCPListener(listener)

	for {
		// Listen for an incoming connection.
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleTCPRequest(conn)
	}
}

func handleTCPRequest(conn *net.TCPConn) {
	defer closeTCPConn(conn)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if _, err := conn.Write([]byte(scanner.Text())); err != nil {
			fmt.Printf("Error writing to connection: %s", err.Error())
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning: %s", err.Error())
	}
}

func udpServerCommand(c *cli.Context) {
	udpAddr, err := net.ResolveUDPAddr("udp", DEFAULT_SERVER_ADDR)
	if err != nil {
		println("Resolving address failed: ", err.Error())
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		println("Dial failed: ", err.Error())
		os.Exit(1)
	}

	defer closeUDPConn(conn)

	buf := make([]byte, len(DEFAULT_BYTES))

	for {
		_, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("Error reading: %s", err.Error())
		}

		replyConn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: addr.IP, Port: DEFAULT_REPLY_PORT_INT})
		if err != nil {
			println("Dial failed: ", err.Error())
			os.Exit(1)
		}

		_, err = replyConn.Write(DEFAULT_BYTES)
		if err != nil {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}

		closeUDPConn(replyConn)
	}
}
