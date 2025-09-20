package main

import "errors"

type LogoutAllCmd struct{}

func (c *LogoutAllCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
