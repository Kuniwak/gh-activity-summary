package main

import (
	"github.com/Kuniwak/gh-activity-summary/cli"
	"github.com/Kuniwak/gh-activity-summary/cmd"
)

func main() {
	cli.Run(cmd.MainCommandByArgs)
}