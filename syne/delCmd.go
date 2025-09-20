package main

type DelCmd struct {
	Path string `arg:"" help:"File or Folder to delete"`
}

func (c *DelCmd) Run() error {
	return nil
}
