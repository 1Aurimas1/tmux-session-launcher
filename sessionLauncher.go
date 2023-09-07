package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

const programName string = "sessionLauncher"

type Config struct {
	ProjectDir         string
	DbStartCommand     string
	GitBranch          string
	SessionName        string
	ServerStartCommand string
	ClientStartCommand string
}

func createNewSession(sessionName string) {
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

func sendKeys(sessionName string, command string) {
	_ = exec.Command("tmux", "send-keys", "-t", sessionName, command, "Enter").Run()
}

func selectWindow(sessionName string) {
	_ = exec.Command("tmux", "select-window", "-t", fmt.Sprintf("%v:2", sessionName)).Run()
}

func splitWindowHorizontally() {
	_ = exec.Command("tmux", "split-window", "-h").Run()
}

func attachSession(sessionName string) {
    err := exec.Command("tmux", "attach-session", "-t", sessionName).Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("started")

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <argument>\n", programName)
        os.Exit(1)
    }

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file")
		os.Exit(1)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Failed to parse file")
		os.Exit(1)
	}

    createNewSession(config.SessionName)

    sendKeys(config.SessionName, "nvim .")

	//if down start /etc/init.d/postgresql status
    createNewWindow()
    selectWindow(config.SessionName)
    initBackend := fmt.Sprintf("cd backend && %v", config.ServerStartCommand)
    sendKeys(config.SessionName, initBackend)
    splitWindowHorizontally()
    initFrontend := fmt.Sprintf("cd frontend && %v", config.ClientStartCommand)
    sendKeys(config.SessionName, initFrontend)

	// duplicate previous window?
    createNewWindow()
    sendKeys(config.SessionName, "cd backend")
    createNewWindow()
    sendKeys(config.SessionName, "cd frontend")

    attachSession(config.SessionName)

	fmt.Println("program exiting")
}
