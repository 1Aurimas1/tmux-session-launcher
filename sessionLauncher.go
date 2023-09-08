package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const programName string = "sessionLauncher"
var sessionName string

type Window struct {
	Commands []string
}

type Config struct {
	ProjectDir  string
	DbStartCmd  string
	SessionName string
	Windows     []Window
}

func startDbServer(dbStartCmd string) {
	if len(dbStartCmd) > 0 {
		_ = exec.Command("bash", "-c", dbStartCmd).Run()
	}
}

func createNewSession() {
	cmd := exec.Command("tmux", "new-session", "-d", "-s", sessionName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func createNewWindow() {
	_ = exec.Command("tmux", "new-window").Run()
}

func sendKeys(command string) {
	_ = exec.Command("tmux", "send-keys", "-t", sessionName, command, "Enter").Run()
}

func selectWindow(windowNum uint8) {
	_ = exec.Command("tmux", "select-window", "-t", fmt.Sprintf("%v:%v", sessionName, windowNum)).Run()
}

func splitWindowHorizontally() {
	_ = exec.Command("tmux", "split-window", "-h").Run()
}

func attachSession() {
	cmd := exec.Command("tmux", "attach-session", "-t", sessionName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <argument>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file\n")
		os.Exit(1)
	}

    var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Failed to parse file\n")
		os.Exit(1)
	}

    if len(config.SessionName) == 0 {
		fmt.Printf("Session name is not set in config file\n")
		os.Exit(1)
    }
    sessionName = config.SessionName

	startDbServer(config.DbStartCmd)

	createNewSession()

	for i, w := range config.Windows {
		if i > 0 {
			createNewWindow()
		}

		for j, cmd := range w.Commands {
			if j > 0 {
				splitWindowHorizontally()
			}

			sendKeys(cmd)
		}
	}

	selectWindow(1)

	attachSession()
}
