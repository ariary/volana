package main

import (
	"volana/pkg/command"
	"volana/pkg/volana"
)

func main() {
	var historic volana.Historic
	for {
		volana.Prefix()
		cmd := volana.GetCommandInteractive(&historic)
		command.ManageSpecialCase(cmd)
		command.Exec(cmd)
	}
}
