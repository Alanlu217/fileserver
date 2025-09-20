package main

type SyncCmd struct {
	Path string `arg:"" help:"File or Folder to upload" default:"~"`
}

func (c *SyncCmd) Run() error {
	return nil
}
