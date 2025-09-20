package main

import "errors"

type SyncCmd struct {
	Path string `arg:"" help:"File or Folder to upload" default:"~"`
}

func (c *SyncCmd) Run() error {
	return errors.New("Not Implemented Yet")
}
