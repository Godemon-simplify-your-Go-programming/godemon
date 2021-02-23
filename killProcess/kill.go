package killProcess

import (
	"godemon/errors"
	"godemon/prepareProject"
	"os"
	"os/exec"
)

func killCMDgen(hOS string) *exec.Cmd {
	name := prepareProject.LoadProjectInfo().Name
	var cmd *exec.Cmd
	if hOS == "windows" {
		cmd = exec.Command("taskkill", "/IM", "app-godemon-app-godemon-tmp-generated"+"-"+name+".exe", "/F")
	} else {
		cmd = exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated"+"-"+name)
	}
	return cmd
}

func KillProcess(hOS string) {
	cmd := killCMDgen(hOS)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errors.TMPerrorHandle(err)
}
