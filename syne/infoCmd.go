package main

type InfoCmd struct {
	Path string `arg:""`
}

func (c *InfoCmd) Run() error {
	return nil
}
