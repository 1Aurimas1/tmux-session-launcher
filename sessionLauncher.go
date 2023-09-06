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

func main() {
	fmt.Printf("started")
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <argument>\n", programName)
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("File not found")
		os.Exit(1)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Failed to parse file")
		os.Exit(1)
	}

	cmd := exec.Command("tmux", "new-session", "-d", "-s", config.SessionName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		panic(err)
	}

	err = exec.Command("tmux", "send-keys", "-t", config.SessionName, "nvim .", "Enter").Run()

	//if down start /etc/init.d/postgresql status
	err = exec.Command("tmux", "new-window").Run()
	err = exec.Command("tmux", "select-window", "-t", fmt.Sprintf("%v:2", config.SessionName)).Run()
	err = exec.Command("tmux", "send-keys", "-t", config.SessionName, fmt.Sprintf("cd backend && %v", config.ServerStartCommand)).Run()
	err = exec.Command("tmux", "split-window", "-h").Run()
	err = exec.Command("tmux", "send-keys", "-t", config.SessionName, fmt.Sprintf("cd frontend && %v", config.ClientStartCommand)).Run()

	// duplicate previous window?
	err = exec.Command("tmux", "new-window").Run()
	err = exec.Command("tmux", "send-keys", "-t", config.SessionName, "cd backend", "Enter").Run()
	err = exec.Command("tmux", "new-window").Run()
	err = exec.Command("tmux", "send-keys", "-t", config.SessionName, "cd frontend", "Enter").Run()

	err = exec.Command("tmux", "attach-session", "-t", config.SessionName).Run()

	fmt.Println("program exiting")
}
