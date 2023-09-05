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

	exec.Command(fmt.Sprintf("cd %v", config.ProjectDir))
    fmt.Println(fmt.Sprintf("cd %v", config.ProjectDir))
	//exec.Command(config.DbStartCommand)
	////exec.Command(fmt.Sprintf("git checkout %v", config.GitBranch))
	//exec.Command(fmt.Sprintf("tmux new-session -d -s %v", config.SessionName))

	//exec.Command(fmt.Sprintf("tmux send-keys -t %v \"nvim .\" Enter", config.SessionName))

	//exec.Command("tmux new-window")
	//exec.Command(fmt.Sprintf("tmux select-window -t %v:2", config.SessionName))
	//exec.Command(fmt.Sprintf("tmux send-keys -t %v \"cd backend && %v\" Enter", config.SessionName, config.ServerStartCommand))
	//exec.Command("tmux split-window -h")
	//exec.Command(fmt.Sprintf("tmux send-keys -t %v \"cd frontend && %v\" Enter", config.SessionName, config.ClientStartCommand))

	//// duplicate previous window?
	//exec.Command("tmux new-window")
	//exec.Command(fmt.Sprintf("tmux send-keys -t %v \"cd backend\" Enter", config.SessionName))
	//exec.Command("tmux new-window")
	//exec.Command(fmt.Sprintf("tmux send-keys -t %v \"cd frontend\" Enter", config.SessionName))

	//exec.Command(fmt.Sprintf("tmux attach-session -t %v", config.SessionName))

    fmt.Println("program exiting")
}
