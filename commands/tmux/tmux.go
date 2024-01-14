package tmux

import (
	"fmt"
	"os"
	"os/exec"
)

const commandName = "tmux"

var SessionName string

func CreateNewSession() {
	cmd := exec.Command(commandName, "new-session", "-d", "-s", SessionName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func CreateNewWindow() {
	_ = exec.Command(commandName, "new-window").Run()
}

func SendKeys(command string) {
	_ = exec.Command(commandName, "send-keys", "-t", SessionName, command, "Enter").Run()
}

func SelectWindow(windowNum uint8) {
	_ = exec.Command(commandName, "select-window", "-t", fmt.Sprintf("%v:%v", SessionName, windowNum)).Run()
}

func SplitWindowHorizontally() {
	_ = exec.Command(commandName, "split-window", "-h").Run()
}

func AttachSession() {
	cmd := exec.Command(commandName, "attach-session", "-t", SessionName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
}
