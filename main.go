package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// take the string from the terminal
	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n') // Reads until Enter is pressed
	cmd = strings.TrimSpace(cmd)
	// store the string as a slice, using a delimiter " "
	args := strings.Split(cmd, " ")
	// validating the command
	err := validateCommand(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	// figuring out the command type
	cmdType := commandType(args)
	// based on the type of command we will perform the necessary actions
	fmt.Println(cmdType)
	if cmdType == 0 {
		fmt.Println("clear cache command")
	} else {
		fmt.Println("starting the proxy command")
	}
}
