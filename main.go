package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var portNum string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// take the string from the terminal
		cmd, _ := reader.ReadString('\n') // Reads until Enter is pressed
		cmd = strings.TrimSpace(cmd)
		// store the string as a slice, using a delimiter " "
		args := strings.Split(cmd, " ")
		if args[0] == "exit" {
			return
		}
		// validating the command
		err := validateCommand(args)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// figuring out the command type
		cmdType := commandType(args)
		// based on the type of command we will perform the necessary actions
		if cmdType == 0 {
			restartServer(":" + portNum)
			fmt.Println("clear cache command")
		} else {
			portNum = args[2]
			startServer(":" + portNum)
			fmt.Println("starting the proxy command")
		}
	}
}
