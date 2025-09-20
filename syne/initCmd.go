package main

import "errors"

type InitCmd struct {
	Path string `arg:"" help:"Subdirectory on server" default:"/"`
}

func (c *InitCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
