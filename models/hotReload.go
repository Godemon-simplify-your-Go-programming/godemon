package models

import "os/exec"

func CMDhotReload(hostInfo [2]string) *exec.Cmd {
	if hostInfo[0] != "windows" {
		cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated")
		return cmd
	} else {
		cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated.exe")
		return cmd
	}
}
