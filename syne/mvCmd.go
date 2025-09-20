package main

import "errors"

type MvCmd struct {
	From string `arg:"" help:"The File / Folder to rename"`
	To   string `arg:"" help:"The new name / path"`
}

func (c *MvCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
