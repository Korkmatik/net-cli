package main

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func NsAction(c *cli.Context) error {
	ns, err := net.LookupNS(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}
	return nil
}

func IPAction(c *cli.Context) error {
	ip, err := net.LookupIP(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}
	return nil
}

func CnameAction(c *cli.Context) error {
	cname, err := net.LookupCNAME(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cname)
	return nil
}

func MxAction(c *cli.Context) error {
	mx, err := net.LookupMX(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i].Host, mx[i].Pref)
	}
	return nil
}
