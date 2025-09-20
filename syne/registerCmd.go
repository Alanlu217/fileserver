package main

import (
	"fmt"
)

type RegisterCmd struct {
	Name string `arg:"" help:"Give a name for the server"`
	Url  string `arg:"" help:"Url to the server"`
}

func (c *RegisterCmd) Run() error {
	fmt.Printf("Adding %s %s\n", Cli.Register.Name, Cli.Register.Url)

	Reg[Cli.Register.Name] = &Server{Url: Cli.Register.Url}
	return nil
}
