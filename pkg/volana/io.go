package volana

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

const prefix = "volana » "

// Get the command and args from user input
func GetCommand() (cmd string) {
	for {
		buf := bufio.NewReader(os.Stdin)
		c, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		cmd += strings.Trim(c, "\n")
		if !strings.HasSuffix(cmd, "\\") {
			return cmd
		} else {
			//Multiple-line command
			cmd = strings.Trim(cmd, "\\")
		}
	}
}

func WriteCharAtIndex(str string, c rune, index int) (new string, err error) {
	sz := len(str)
	if index == sz {
		new = str + string(c)

	} else if index < sz {
		new = str[:index] + string(c) + str[index:]
	}

	return new, fmt.Errorf("WriteCharAtIndex: Index out of range")
}

//Return the last character of a string.
func GetLastCharacter(str string) string {
	sz := len(str)
	if sz > 0 {
		return str[sz-1:]
	}
	return ""
}

//Return the first character of a string.
func GetFirstCharacter(str string) string {
	sz := len(str)
	if sz > 0 {
		return str[:1]
	}
	return ""
}

// Get the command and args from user input + interactive behiviour ie react with arrow key stroke, backspace
//TODO: handle delete/suppr key
func GetCommandInteractive(historic *Historic) (cmd string) {
	//keyboard listener
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	//For interativeness with lateral movement with arrow
	var cmdRight string

	//For multipleline printing
	var multipleline bool
	var line string
	var lineRight string

	for { //while(1)
		event := <-keysEvents
		if event.Err != nil {
			fmt.Println(event.Err)
		}
		switch event.Key {
		case keyboard.KeyEnter:
			//Enter: Validate command Or continue if it ends w/ '\'
			if !strings.HasSuffix(cmd, "\\") { //check if cmdline does not finish with '\'
				fmt.Println()
				fullCmd := cmd + cmdRight
				historic.Add(fullCmd)
				return fullCmd
			} else {
				//Multiple-line command
				fmt.Println()
				cmd = strings.Trim(cmd, "\\")
				multipleline = true
			}
		case keyboard.KeySpace:
			//event.Rune for space could be x00 => error
			space := " "
			fmt.Print(space)
			cmd += space
		case keyboard.KeyBackspace:
			//delete one character
			sz := len(cmd)
			if sz > 2 {
				cmd = cmd[:sz-1]
				if multipleline {
					fmt.Printf("\r%s ", line)
				} else {
					fmt.Printf("\r%s%s ", prefix, cmd)
				}
			}
		case keyboard.KeyArrowUp:
			//get previous command
			previous := historic.GetPrevious()
			if previous != "" {
				cmd = previous
				cmdRight = ""
				fmt.Printf("\r%s%s", prefix, previous)
			}
		case keyboard.KeyArrowDown:
			next := historic.GetNext()
			if next != "" {
				cmd = next
				cmdRight = ""
				fmt.Printf("\r%s%s", prefix, next)
			}
		case keyboard.KeyArrowLeft:
			//Shift writing place to the left + adapt printing
			//does not handle multipleline
			last := GetLastCharacter(cmd)
			cmd = strings.TrimSuffix(cmd, last)
			cmdRight = last + cmdRight
			if multipleline {
				line = strings.TrimSuffix(line, last)
				lineRight = last + lineRight
				fmt.Printf("\r%s", line)
			} else {
				fmt.Printf("\r%s%s", prefix, cmd) //shift writing index by one
			}
		case keyboard.KeyArrowRight:
			//Shift writing place to the left + adapt printing
			first := GetFirstCharacter(cmdRight)
			cmdRight = strings.TrimPrefix(cmdRight, first)
			cmd += first
			if multipleline {
				lineRight = strings.TrimPrefix(lineRight, first)
				fmt.Printf("\r%s", line) //shift writing index by ones
			} else {
				fmt.Printf("\r%s%s", prefix, cmd) //shift writing index by ones
			}
		default:
			//Write onecharacter
			if multipleline {
				fmt.Printf("\r%s%s%s", line, string(event.Rune), lineRight)
				line += string(event.Rune)
			} else {
				fmt.Printf("\r%s%s%s%s", prefix, cmd, string(event.Rune), cmdRight)
			}
			cmd += string(event.Rune)
		}
	}
}

//Print the prefix of the shell
func Prefix() {
	fmt.Print(prefix)
}
