package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"tmux-session-launcher/commands/bash"
	"tmux-session-launcher/commands/tmux"
)

const programName = "sessionLauncher"

type Window struct {
	Commands []string
}

type Config struct {
	ProjectDir  string
	DbStartCmd  string
	SessionName string
	Windows     []Window
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
	tmux.SessionName = config.SessionName

	bash.StartDatabase(config.DbStartCmd)

	tmux.CreateNewSession()

	for i, w := range config.Windows {
		if i > 0 {
			tmux.CreateNewWindow()
		}

		for j, cmd := range w.Commands {
			if j > 0 {
				tmux.SplitWindowHorizontally()
			}

			tmux.SendKeys(cmd)
		}
	}

	tmux.SelectWindow(1)

	tmux.AttachSession()
}
