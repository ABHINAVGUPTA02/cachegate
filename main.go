package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var portNum string
var originServer string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// take the string from the terminal
		cmd, _ := reader.ReadString('\n')
		// removing the '\n'
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
			clearCache()
			fmt.Println("clear cache command")
		} else {
			portNum = args[2]
			originServer = args[4]
			r := SetupRouter()
			fmt.Printf("server running on port: %s for origin server: %s\n", portNum, originServer)
			err = http.ListenAndServe(
				":"+portNum,
				r,
			)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("starting the proxy command")
		}
	}
}
