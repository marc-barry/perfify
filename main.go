package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "perfify"
	app.Usage = "Network performance testing."

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "start as server",
			Subcommands: []cli.Command{
				{
					Name:  "tcp",
					Usage: "tcp server",
					Action: func(c *cli.Context) {
						tcpServerCommand(c)
					},
				},
				{
					Name:  "udp",
					Usage: "udp server",
					Action: func(c *cli.Context) {
						udpServerCommand(c)
					},
				},
			},
		},
		{
			Name:  "client",
			Usage: "start as client",
			Subcommands: []cli.Command{
				{
					Name:  "tcp",
					Usage: "tcp related commands",
					Subcommands: []cli.Command{
						{
							Name:  "ping",
							Usage: "perform a tcp ping",
							Action: func(c *cli.Context) {
								tcpPingCommand(c)
							},
							Flags: []cli.Flag{
								cli.IntFlag{
									Name:  N_NAME,
									Value: 1,
									Usage: "the number of ping requests",
								},
								cli.IntFlag{
									Name:  SIZE_NAME,
									Value: 1,
									Usage: "the size of data in bytes",
								},
							},
						},
					},
				},
				{
					Name:  "udp",
					Usage: "udp related commands",
					Subcommands: []cli.Command{
						{
							Name:  "ping",
							Usage: "perform a udp ping",
							Action: func(c *cli.Context) {
								udpPingCommand(c)
							},
							Flags: []cli.Flag{
								cli.IntFlag{
									Name:  N_NAME,
									Value: 1,
									Usage: "the number of ping requests",
								},
								cli.IntFlag{
									Name:  SIZE_NAME,
									Value: 1,
									Usage: "the size of data in bytes",
								},
							},
						},
					},
				},
			},
		},
		{
			Name:  "interfaces",
			Usage: "list interface info",
			Action: func(c *cli.Context) {
				printInterfaces()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		println("Error running CLI app: ", err.Error())
	}
}
