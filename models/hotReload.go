package models

import "os/exec"

func CMDhotReload(hostInfo [2]string, name string) *exec.Cmd {
	if hostInfo[0] != "windows" {
		cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated"+"-"+name)
		return cmd
	} else {
		cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated"+"-"+name+".exe")
		return cmd
	}
}
