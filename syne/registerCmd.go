package main

import (
	"fmt"
)

type RegisterCmd struct {
	Force bool   `short:"f" help:"Don't check if server works"`
	Name  string `short:"n" help:"Give a name for the server" default:""`

	Url string `arg:"" help:"Url to the server"`
}

func (c *RegisterCmd) Run() error {
	if c.Force && c.Name == "" {
		return fmt.Errorf("Register requires name when using force.")
	}

	if c.Name == "" {
		c.Name = "Get from server"
	}

	fmt.Printf("Adding %s %s\n", Cli.Register.Name, Cli.Register.Url)
	Reg[Cli.Register.Name] = &Server{Url: Cli.Register.Url}

	return Reg.Write()
}
