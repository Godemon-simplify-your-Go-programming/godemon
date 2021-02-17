package controllers

import (
	"os"
	"os/exec"
)

func execMOD(hOS string) {
	if hOS == "windows" {
		cmd := exec.Command("app-godemon-app-godemon-tmp-generated.exe")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	} else {
		cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	}
}

func execFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}
