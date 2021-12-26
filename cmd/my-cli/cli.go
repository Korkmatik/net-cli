package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and name Servers!"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ns",
			Usage:  "Looks up the Name Servers for a Particular Host",
			Flags:  myFlags,
			Action: NsAction,
		},
		{
			Name:   "ip",
			Usage:  "Looks up the IP addresses for a particular host",
			Flags:  myFlags,
			Action: IPAction,
		},
		{
			Name:   "cname",
			Usage:  "Looks up the CNAME for a particular host",
			Flags:  myFlags,
			Action: CnameAction,
		},
		{
			Name:   "mx",
			Usage:  "Looks up the MX records for a particular host",
			Flags:  myFlags,
			Action: MxAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
