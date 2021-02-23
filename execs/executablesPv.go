package execs

import (
	"godemon/errors"
	"godemon/prepareProject"
	"os"
	"os/exec"
)

func execMOD(hOS string) {
	name := prepareProject.LoadProjectInfo().Name
	if hOS == "windows" {
		cmd := exec.Command("app-godemon-app-godemon-tmp-generated" + "-" + name + ".exe")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		errors.ErrorHandle(err)
	} else {
		cmd := exec.Command("./app-godemon-app-godemon-tmp-generated" + "-" + name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		errors.ErrorHandle(err)
	}
}

func execFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errors.ErrorHandle(err)
}
