package main

import "errors"

type LoginCmd struct {
}

func (c *LoginCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
