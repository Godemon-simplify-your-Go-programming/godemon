package controllers

import (
	"os"
	"os/exec"
)

func killProcess() {
	cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}

func execMOD() {
	cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}

func execFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}
