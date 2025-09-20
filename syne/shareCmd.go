package main

type ShareCmd struct {
	Path      string `arg:""`
	Password  string `short:"p"`
	NumUses   int    `short:"n"`
	AliveTime int    `short:"t"`
}

func (c *ShareCmd) Run() error {
	return nil
}
