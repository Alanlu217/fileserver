package main

type AddCmd struct {
	From string `arg:"" help:"What file to upload" type:"path"`
	To   string `arg:"" help:"Where to upload the file" default:"/"`
}

func (c *AddCmd) Run() error {
	return nil
}
