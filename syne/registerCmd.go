package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/yarlson/pin"
)

type RegisterCmd struct {
	Force bool   `short:"f" help:"Don't check if server works"`
	Name  string `short:"n" help:"Give a name for the server" default:""`

	Url string `arg:"" help:"Url to the server"`
}

func (c *RegisterCmd) Run() error {
	if c.Force {
		if c.Name == "" {
			return fmt.Errorf("Register requires name when using force.")
		}

		Reg[c.Name] = &Server{Url: c.Url}
		return Reg.Write()
	}

	if c.Name == "" {
		c.Name = "Get from server"
	}

	c.Name = strings.TrimSpace(c.Name)

	path := fmt.Sprintf("http://%s/name", c.Url)

	p := pin.New("Getting Server Name...")
	cancel := p.Start(context.Background())
	defer cancel()

	name, err := http.Get(path)

	if err != nil {
		return err
	}
	if name.StatusCode != 200 {
		return fmt.Errorf("Server error")
	}
	name_bytes, err := io.ReadAll(name.Body)
	if err != nil {
		return err
	}
	c.Name = string(name_bytes)

	fmt.Printf("Adding %s %s\n", c.Name, c.Url)
	Reg[c.Name] = &Server{Url: c.Url}
	p.Stop("Done!")

	return Reg.Write()
}
