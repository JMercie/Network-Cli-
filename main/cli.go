package main

import (
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNames, MX records and Name Servers!"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "joseph.dev",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the name Servers for particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, error := net.LookupNS(c.String("host"))
				if error != nil {
					log.Println(error)
					return error
				}
				for i := 0; i < len(ns); i++ {
					log.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					log.Println(err)
					return err
				}
				for i := 0; i < len(ip); i++ {
					log.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME  for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					log.Println(err)
					return err
				}
				log.Println(cname)

				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the mx  for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					log.Println(err)
					return err
				}
				for i := 0; i < len(mx); i++ {
					log.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
