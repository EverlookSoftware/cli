package main

import (
	"fmt"
	"os"
	"log"
)

var possibleCommands = []string{
	"create",
	"help",
}

func createProject() {
	fmt.Println("Creating...");
}

func displayHelp() {
	fmt.Println("Help...");
}

func executeCmd(cmd string) {
	switch cmd {
		case "create":
			createProject();
		case "help":
			displayHelp();
		default:
			log.Fatal("Command " + cmd + "Does not exist.");
	}
}

func checkIfCommandExists(cmd string) bool {
	for _, c := range possibleCommands {
		if c == cmd {
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

	executeCmd(cmd);
}