package killProcess

import (
	"godemon/prepareProject"
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
