package command

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Manage special case by looking at the command itself
func ManageSpecialCase(cmd string) {
	switch cmd {
//	case "cd":
//		if len(args) < 2 {
//        		return  errors.New("path required")
//    		}
//    		os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
}

//Exec a command and return the output as the string. (return also stderr).
//pipe: if a command failed, it is write to stderr but the command keeps running
//TODO: modify for pipe: case when len > 1 and simple case
func Exec(cmd string) {
	cmds := SplitPipe(cmd)

	var input string
	var output string

	if len(cmds) == 1 {
		args := strings.Fields(cmd)
		command := exec.Command(args[0], args[1:]...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin
		err := command.Run()
		if err != nil {
			log.Fatalf("Run command failed with %s\n", err)
		}
	} else { //pipe command
		for i := 0; i < len(cmds); i++ {
			args := strings.Fields(cmd)
			command := exec.Command(args[0], args[1:]...)
			if i == len(cmds)-1 { //Last command of the pipe -> show output in real time
				command.Stdout = os.Stdout
				command.Stderr = os.Stderr
				command.Stdin = os.Stdin
				err := command.Run()
				if err != nil {
					log.Fatalf("Run command failed with %s\n", err)
				}
			} else { //we are within pipe
				stdin, err := command.StdinPipe()
				if err != nil {
					fmt.Println("Exec:", err)
				}

				go func() { //take precedent output of command as input (pipe)
					defer stdin.Close()
					io.WriteString(stdin, input)
				}()

				out, err2 := command.CombinedOutput()
				output = strings.Trim(string(out), "\n")
				if err2 != nil {
					command.Stderr.Write([]byte(output))
				} else {
					input = output //for pipe
				}
			}
		}
	}

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
