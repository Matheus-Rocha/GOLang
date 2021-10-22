package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Kakashi"
	app.Usage = "Search IPs and Server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search public IPS",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "server",
			Usage:  "Search server names",
			Flags:  flags,
			Action: seachServerName,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Println("Look up IP error:", err)
	}

	for _, ip := range ips {
		fmt.Println("Public IP found: ", ip)
	}
}

func seachServerName(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Println("Look up NS error:", err)
	}

	for _, server := range servers {
		fmt.Println("Server: ", server.Host)
	}
}
