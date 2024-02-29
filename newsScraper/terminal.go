
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clearTerminal() {
	cmd := exec.Command("clear") // Use "cls" on Windows.
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing terminal:", err)
	}
}

func displayOptions() {
	fmt.Println("Which Option would you like to use?")
	fmt.Println("[1] - Paste in a fixed URL from an Article")
	fmt.Println("[2] - Get the Latest Article")
}

func readUserInput() (string, error) {
	var userInput string
	_, err := fmt.Scanln(&userInput)
	return userInput, err
}