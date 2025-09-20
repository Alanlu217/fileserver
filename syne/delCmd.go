package main

import "errors"

type DelCmd struct {
	Path string `arg:"" help:"File or Folder to delete"`
}

func (c *DelCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
