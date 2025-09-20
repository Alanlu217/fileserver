package main

import "errors"

type GetCmd struct {
	Path string `arg:"" help:"Path to get" default:"/"`
}

func (c *GetCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
