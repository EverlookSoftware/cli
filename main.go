package main

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"github.com/ttacon/chalk"
)

var help = `
--------------------------------------
[help] - Prompts command list
[create] - bootstraps a new project
--------------------------------------
`;

var CLIENT_DIR_NAME = "client";

type fn func ()

func createClient(name string, c chan string) {
	clone := exec.Command("git", "clone", "https://github.com/EverlookSoftware/react-boilerplate.git", CLIENT_DIR_NAME);
	cloneErr := clone.Run();

	if cloneErr != nil {
		log.Fatal("An error occurred while cloning UI into ", name);
	}

	copy := exec.Command("cp", "-a", CLIENT_DIR_NAME, name);
	copy.Run();
	rm := exec.Command("rm", "-rf", CLIENT_DIR_NAME);
	rm.Run();

	c <- "Frontend successfully created";
}

func createProject() {
	name := os.Args[2];
	fmt.Println(chalk.Magenta.Color("Bootstrapping projects..."));

	mkDir := exec.Command("mkdir", name);
	mkErr := mkDir.Run();

	if mkErr != nil {
		log.Fatal("An error occurred while creating directory ", name);
	}

	c := make(chan string);
	go createClient(name, c);
	// go createServer(name, c);
	
	fmt.Println(chalk.Blue.Color(<- c));
}

func displayHelp() {
	fmt.Println(help);
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