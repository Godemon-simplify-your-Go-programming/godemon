package killProcess

import (
	"godemon/errors"
	"os"
	"os/exec"
)

func killCMDgen(hOS string) *exec.Cmd {
	var cmd *exec.Cmd
	if hOS == "windows" {
		cmd = exec.Command("taskkill", "/IM", "app-godemon-app-godemon-tmp-generated.exe", "/F")
	} else {
		cmd = exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
	}
	return cmd
}

func KillProcess(hOS string) {
	cmd := killCMDgen(hOS)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errors.ErrorHandle(err)
}
