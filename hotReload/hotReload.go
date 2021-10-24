package hotReload

import (
	"godemon/prepareProject"
	"os"
	"os/exec"
)

func CMDhotReload(hostInfo [2]string) *exec.Cmd {
	name := prepareProject.LoadProjectInfo().Name
	if hostInfo[0] != "windows" {
		cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated"+"-"+name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd
	} else {
		cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated"+"-"+name+".exe")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd
	}
}
