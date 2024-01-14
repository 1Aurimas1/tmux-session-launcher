package bash

import "os/exec"

const bashCommandName = "bash"

func StartDatabase(dbStartCmd string) {
	if len(dbStartCmd) > 0 {
		_ = exec.Command(bashCommandName, "-c", dbStartCmd).Run()
	}
}
