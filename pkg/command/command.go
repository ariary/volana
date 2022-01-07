package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Manage special case by looking at the command itself
func ManageSpecialCase(cmd string) {
	switch cmd {
	case "exit":
		os.Exit(0)
	}
}

//Exec a command and return the output as the string. (return also stderr).
//pipe: if a command failed, it is write to stderr but the command keeps running
func Exec(cmd string) (result string) {
	cmds := SplitPipe(cmd)

	var stderr string
	var input string
	var output string

	for i := 0; i < len(cmds); i++ {
		shell := exec.Command("/bin/sh", "-c", cmd)
		stdin, err := shell.StdinPipe()
		if err != nil {
			fmt.Println("Exec:", err)
		}

		go func() { //take precedent output of command as input (pipe)
			defer stdin.Close()
			io.WriteString(stdin, input)
		}()

		out, err2 := shell.CombinedOutput()
		output = strings.Trim(string(out), "\n")
		if err2 != nil {
			stderr += string(output)
		} else {
			input = output //for pipe
		}
	}
	if stderr != "" {
		output += output + "\n" + stderr
	}
	return output
}

//Take a command in input and return a slice of each command separated in order, if there is a pipe.
//It returns the command in a slice of lenght 1 if there isn't any pipe
func SplitPipe(command string) (commands []string) {
	//regex matching the command before a pipe
	re1, err := regexp.Compile(`^([^"]|"[^"]*")*?(\|)`)
	if err != nil {
		fmt.Println("regexp:", err)
	}

	//Split each commands
	for {
		var c string
		if re1.MatchString(command) {
			//pipe
			c = strings.TrimSuffix(re1.FindString(command), "|")
			commands = append(commands, c)
			command = re1.ReplaceAllString(command, "")
		} else {
			//no more pipe
			commands = append(commands, command)
			break
		}
	}
	return commands
}
