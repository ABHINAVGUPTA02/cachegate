package main

import (
	"errors"
	"strconv"
)

func commandType(args []string) int {
	if len(args) == 2 {
		return 0
	}
	return 1
}

func validateCommand(args []string) error {
	if len(args) < 2 {
		return errors.New("not enough arguments")
	}

	if len(args) == 2 {
		if string(args[0]) != "caching-proxy" && string(args[1]) != "--clear-cache" {
			return errors.New("unsupported command")
		}

		return nil
	}

	if len(args) == 5 {
		if string(args[0]) != "caching-proxy" && string(args[1]) != "--port" && string(args[3]) != "--origin" {
			return errors.New("unsupported command")
		}

		portNumber, _ := strconv.Atoi(args[2])

		if portNumber < 1 || portNumber > 65535 {
			return errors.New("invalid port number")
		}

		return nil
	}

	return errors.New("invalid command")
}
