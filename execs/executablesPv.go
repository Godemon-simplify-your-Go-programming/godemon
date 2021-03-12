package execs

import (
	"godemon/errors"
	"godemon/models"
	"godemon/prepareProject"
	"os"
	"os/exec"
	"strings"
)

func execMOD(hOS string, flags []models.Flag) {
	name := prepareProject.LoadProjectInfo().Name
	var flagsF []string
	i := 0
	for i < len(flags) {
		flagsF = append(flagsF, flags[i].Flag)
		i++
	}
	f := strings.Join(flagsF, " ")
	if hOS == "windows" {
		cmd := exec.Command("app-godemon-app-godemon-tmp-generated"+"-"+name+".exe", f)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		errors.ErrorHandle(err)
	} else {
		cmd := exec.Command("./app-godemon-app-godemon-tmp-generated"+"-"+name, f)
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
