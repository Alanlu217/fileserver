package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
)

var Reg ServerRegistry

var Cli struct {
	Guest  bool   `short:"g" help:"Execute as guest"`
	Server string `short:"s" help:"Server to use"`

	Register RegisterCmd `cmd:"" help:"Registers a server"`

	Login     LoginCmd     `cmd:"" help:"Login to account"`
	Logout    LogoutCmd    `cmd:"" help:"Logout from account"`
	Logoutall LogoutAllCmd `cmd:"" help:"Logout from all accounts"`

	Init InitCmd `cmd:"" help:"Initialises a directory for syncing"`
	Sync SyncCmd `cmd:"" help:"Syncs the active folder"`

	Get  GetCmd  `cmd:"" help:"Download a file"`
	Add  AddCmd  `cmd:"" help:"Add a File or Folder to the server"`
	Mv   MvCmd   `cmd:"" help:"Renames a File or Folder on the server"`
	Del  DelCmd  `cmd:"" help:"Deletes a File or Folder"`
	Info InfoCmd `cmd:"" help:"Get info on a file or folder"`

	Share ShareCmd `cmd:"" help:"Temporarily share a file or folder"`
}

func main() {
	var err error
	ctx := kong.Parse(&Cli)

	Reg, err = ParseRegistry()
	if err != nil {
		log.Fatal(err)
	}

	err = ctx.Run()
	if err != nil {
		fmt.Println(err)
	}
}
