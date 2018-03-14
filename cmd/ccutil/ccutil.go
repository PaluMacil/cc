package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PaluMacil/cc/action"
)

const (
	Usage = `

	Usage:
				ccutil [cmd] [arg] [arg2]

	Commands:
		add		[name] [path]
				adds a compiler name and path

		remove	[name]
				removes a compiler by name

		select	[name]
				selects a compiler to be the default

		list
				lists compiler names and paths

		status
				prints the active compiler

	`
)

func main() {
	// Run with test argument to type arguments after launching (helps debugging)
	if len(os.Args) == 2 && (os.Args[1] == "test" || os.Args[1] == "-t") {
		fmt.Println("Test mode. Please enter args:")
		reader := bufio.NewReader(os.Stdin)
		args, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Invalid args.")
		}
		os.Args = append(os.Args[:1], strings.Split(args, " ")...)
	}
	if len(os.Args) == 2 && (os.Args[1] == "help" || os.Args[1] == "-h") {
		log.Println(Usage)
		return
	}
	a, err := action.FromOSArgs(os.Args)
	if err != nil {
		log.Fatalf("could not create action: %s from os.Args: %v", err, os.Args)
	}
	err = a.Execute(os.Stderr)
	if err != nil {
		log.Fatalf("could not execute action: %s from os.Args: %v", err, os.Args)
	}
}
