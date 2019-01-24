package main

import (
	"fmt"
	"os"
	"log"
)

type fn func ()

func createProject() {
	name := os.Args[2];
	
	fmt.Println("Creating...", name);
}

func displayHelp() {
	fmt.Println("Help...");
}

var commands = map[string] fn {
	"create": createProject,
	"help": displayHelp,
}

func checkIfCommandExists(cmd string) bool {
	for key := range commands {
		if key == cmd {
			return true;
		}
	}

	return false;
}

func main() {
	cmd := os.Args[1];
	doesCmdExist := checkIfCommandExists(cmd);

	if doesCmdExist == false {
		log.Fatal("Command " + cmd + "Does not exist.");
	}

	commands[cmd]();
}