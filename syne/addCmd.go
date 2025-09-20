package main

import "errors"

type AddCmd struct {
	From string `arg:"" help:"What file to upload" type:"path"`
	To   string `arg:"" help:"Where to upload the file" default:"/"`
}

func (c *AddCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
