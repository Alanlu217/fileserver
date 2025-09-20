package main

import "errors"

type InfoCmd struct {
	Path string `arg:""`
}

func (c *InfoCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
