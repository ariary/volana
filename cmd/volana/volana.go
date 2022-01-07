package main

import (
	"fmt"
	"volana/pkg/command"
	"volana/pkg/volana"
)

func main() {
	var historic volana.Historic
	for {
		volana.Prefix()
		cmd := volana.GetCommandInteractive(&historic)
		command.ManageSpecialCase(cmd)
		output := command.Exec(cmd)
		fmt.Println(output)
	}
}
