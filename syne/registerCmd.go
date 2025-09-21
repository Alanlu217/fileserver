package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

		Pin.Stop("Done!")
		return Reg.Write()
	}

	c.Name = strings.TrimSpace(c.Name)

	if c.Name == "" {
		path := fmt.Sprintf("http://%s/name", c.Url)

		Pin.UpdateMessage("Getting Server Name...")
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
	}

	fmt.Printf("Adding %s %s\n", c.Name, c.Url)
	Reg[c.Name] = &Server{Url: c.Url}

	Pin.Stop("Done!")
	return Reg.Write()
}
