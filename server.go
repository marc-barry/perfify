package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"net"
	"os"
	"strconv"
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
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	sizeStr, err := rw.ReadString('\n')
	if err != nil {
		fmt.Printf("Can't read size: %s", err.Error())
		return
	}

	sizeStr = sizeStr[0 : len(sizeStr)-1]

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		fmt.Printf("Can't convert size to int: %s", err.Error())
		return
	}

	fmt.Printf("Pinged with %s bytes.\n", sizeStr)

	buf := make([]byte, 1024)

	totalRead := 0
	totalWrote := 0

	for {
		nr, err := rw.Read(buf)
		if err != nil || nr == 0 {
			break
		}

		totalRead += nr

		nw, err := rw.Write(buf[0:nr])
		if err != nil {
			break
		}

		err = rw.Flush()
		if err != nil {
			fmt.Printf("Error fluhing buffer: %s", err.Error())
		}

		totalWrote += nw

		if totalRead == size {
			break
		}
	}

	fmt.Printf("Read %d bytes.\n", totalRead)
	fmt.Printf("Wrote %d bytes.\n", totalWrote)
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

		_, err = conn.WriteToUDP(DEFAULT_BYTES, addr)
		if err != nil {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}
	}
}
