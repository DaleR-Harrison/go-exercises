package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Print("What is your name?: ")
	fmt.Println(determineMsg(handleReadInString(bufio.NewReader(os.Stdin).ReadString('\n'))))
}

func handleReadInString(rawString string, ioErr error) (name string, err error) {
	if err != nil {
		return "", ioErr
	}

	return strings.Trim(rawString, "\n"), nil
}

func determineMsg(name string, err error) string {
	if err != nil {
		return "I'm sorry, I don't understand."
	}

	switch {
	case name == "Dale":
		return "Hey, we have the same name! Hi, other Dale."
	case name == "Hikari":
		return "Hello, my love. How are you?"
	default:
		return fmt.Sprintf("Hello, %s", name)
	}
}
