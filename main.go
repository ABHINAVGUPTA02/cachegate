package main

import (
	"bufio"
	"github.com/fatih/color"
	"net/http"
	"os"
	"strings"
)

var portNum string
var originServer string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		args := strings.Split(cmd, " ")
		if args[0] == "exit" {
			return
		}
		err := validateCommand(args)
		if err != nil {
			color.Red(err.Error())
			continue
		}
		cmdType := commandType(args)

		if cmdType == 0 {
			clearCache()
			color.Green("Cache cleared...")
			continue
		} else {
			portNum = args[2]
			originServer = args[4]
			r := SetupRouter()

			go func() {
				http.ListenAndServe(
					":"+portNum,
					r,
				)
			}()
			color.Green("Server started...")
		}
	}
}
