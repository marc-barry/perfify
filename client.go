package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"net"
	"os"
	"time"
)

func tcpPingCommand(c *cli.Context) {
	n := c.Int(N)

	fmt.Printf("Ping TCP endpoint %d times.\n", n)

	tcpAddr, err := net.ResolveTCPAddr("tcp", DEFAULT_SERVER_ADDR)
	if err != nil {
		println("Resolving address failed: ", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed: ", err.Error())
		os.Exit(1)
	}

	defer closeTCPConn(conn)

	var totalRtt int64 = 0
	var totalSamples int64 = int64(n)

	for i := 1; i <= n; i++ {
		t := time.Now()

		_, err = conn.Write(DEFAULT_BYTES)
		if err != nil {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}

		reply := make([]byte, len(DEFAULT_BYTES))

		_, err = conn.Read(reply)
		if err != nil {
			println("Read from server failed: ", err.Error())
			os.Exit(1)
		}

		rtt := time.Since(t)

		fmt.Printf("TCP Ping(%d): %s\n", i, rtt.String())

		totalRtt += rtt.Nanoseconds()
	}

	fmt.Printf("Average TCP Ping: %s\n", time.Duration(totalRtt/totalSamples).String())
}

func udpPingCommand(c *cli.Context) {
	n := c.Int(N)

	fmt.Printf("Ping UDP endpoint %d times.\n", n)

	replyAddr, err := net.ResolveUDPAddr("udp", DEFAULT_REPLY_ADDR)
	if err != nil {
		println("Resolving address failed: ", err.Error())
		os.Exit(1)
	}

	replyConn, err := net.ListenUDP("udp", replyAddr)
	if err != nil {
		println("Dial failed: ", err.Error())
		os.Exit(1)
	}

	serverAddr, err := net.ResolveUDPAddr("udp", DEFAULT_SERVER_ADDR)
	if err != nil {
		println("Resolving address failed: ", err.Error())
		os.Exit(1)
	}

	defer closeUDPConn(replyConn)

	var totalRtt int64 = 0
	var totalSamples int64 = int64(n)

	for i := 1; i <= n; i++ {
		t := time.Now()

		_, err = replyConn.WriteToUDP(DEFAULT_BYTES, serverAddr)
		if err != nil {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}

		reply := make([]byte, len(DEFAULT_BYTES))

		_, _, err := replyConn.ReadFromUDP(reply)
		if err != nil {
			println("Read from server failed: ", err.Error())
			os.Exit(1)
		}

		rtt := time.Since(t)

		fmt.Printf("UDP Ping(%d): %s\n", i, rtt.String())

		totalRtt += rtt.Nanoseconds()
	}

	fmt.Printf("Average UDP Ping: %s\n", time.Duration(totalRtt/totalSamples).String())
}
