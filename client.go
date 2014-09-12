package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/codegangsta/cli"
	"net"
	"os"
	"strconv"
	"time"
)

func tcpPingCommand(c *cli.Context) {
	n := c.Int(N_NAME)
	size := c.Int(SIZE_NAME)
	size_str := strconv.Itoa(size)

	data := make([]byte, size)
	count, err := rand.Read(data)
	if err != nil {
		println("Error genarating random data: ", err.Error())
		os.Exit(1)
	}
	if count != size {
		println("Size mismatch.")
		os.Exit(1)
	}

	fmt.Printf("Ping TCP endpoint %d times with %d bytes.\n", n, size)

	var totalRtt int64 = 0
	var totalSamples int64 = int64(n)

	for i := 1; i <= n; i++ {
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

		t := time.Now()

		_, err = conn.Write(append([]byte(size_str), '\n'))
		if err != nil {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}

		nw, err := conn.Write(data)
		if err != nil || nw != len(data) {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}

		buf := make([]byte, 1000)

		total := 0

		r := bufio.NewReader(conn)

		for {
			nr, err := r.Read(buf)
			if err != nil && nr == 0 {
				break
			}

			total += nr

			if total == size {
				break
			}
		}

		rtt := time.Since(t)

		if total != size {
			fmt.Printf("The received data is a different size (sent: %d, recieved: %d).\n", nw, total)
		}

		fmt.Printf("TCP Ping(%d): %s\n", i, rtt.String())

		totalRtt += rtt.Nanoseconds()

		closeTCPConn(conn)
	}

	fmt.Printf("Average TCP Ping: %s\n", time.Duration(totalRtt/totalSamples).String())
}

func udpPingCommand(c *cli.Context) {
	n := c.Int(N_NAME)
	size := c.Int(SIZE_NAME)
	data := make([]byte, size)
	_, err := rand.Read(data)

	fmt.Printf("Ping UDP endpoint %d times with %d bytes.\n", n, size)

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

		_, err = replyConn.WriteToUDP(data, serverAddr)
		if err != nil {
			println("Write to server failed: ", err.Error())
			os.Exit(1)
		}

		reply := make([]byte, len(data))

		_, _, err := replyConn.ReadFromUDP(reply)
		if err != nil {
			println("Read from server failed: ", err.Error())
			os.Exit(1)
		}

		rtt := time.Since(t)

		if !bytes.Equal(data, reply) {
			fmt.Printf("The sent and received data is not equal.\n")
		}

		fmt.Printf("UDP Ping(%d): %s\n", i, rtt.String())

		totalRtt += rtt.Nanoseconds()
	}

	fmt.Printf("Average UDP Ping: %s\n", time.Duration(totalRtt/totalSamples).String())
}
