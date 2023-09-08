package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns command line app
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command line app"
	app.Usage = "Search IPs and servers name in Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search address IPs in Internet",
			Flags:  flags,
			Action: SearchIp,
		},
		{
			Name:   "server",
			Usage:  "Search servers name in Internet",
			Flags:  flags,
			Action: SearchServer,
		},
	}

	return app
}

func SearchIp(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func SearchServer(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
