package execs

import (
	"godemon/errors"
	"os"
	"os/exec"
)

func execMOD(hOS string, name string, logs bool) {
	if hOS == "windows" {
		cmd := exec.Command("app-godemon-app-godemon-tmp-generated" + "-" + name + ".exe")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		errors.ErrorHandle(err)
	} else {
		if logs == true {
			logFile := name + "_logs.txt"
			cmd := exec.Command("./app-godemon-app-godemon-tmp-generated" + "-" + name + "2>&1" + "|" + "tee" + logFile)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			errors.ErrorHandle(err)
			return
		}
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
