package main

import "errors"

type LogoutCmd struct {
	Others bool `short:"o" help:"Logout all others"`
}

func (c *LogoutCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
