package killProcess

import (
	"godemon/errors"
	"os"
	"os/exec"
)

func killCMDgen(hOS string, name string) *exec.Cmd {
	var cmd *exec.Cmd
	if hOS == "windows" {
		cmd = exec.Command("taskkill", "/IM", "app-godemon-app-godemon-tmp-generated"+"-"+name+".exe", "/F")
	} else {
		cmd = exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated"+"-"+name)
	}
	return cmd
}

func KillProcess(hOS string, name string) {
	cmd := killCMDgen(hOS, name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errors.TMPerrorHandle(err)
}
