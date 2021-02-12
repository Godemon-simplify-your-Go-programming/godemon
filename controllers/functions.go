package controllers

import (
	"godemon/models"
	"os"
	"time"
)

func ExecMOD() {
	execMOD()
}

func ExecFile(filepath string) {
	execFile(filepath)
}

func TimeLog() {
	timeLog()
}

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, version string, init bool, name string, oso string, arch string) (string, string) {
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile = cnfFunc(command, "", "")
	} else if *cnf == "deploy" {
		deploy()
	} else if init == true {
		initialize(name, arch, oso)
	} else if *help == true ||
		(*cnf == "" && filepath == "" && modOrFile == "" &&
			command == "" && *help == false && init == false &&
			name == "" && oso == "" && arch == "") {
		models.HelpCLI(version)
	}
	return filepath, modOrFile
}

func WatchFiles(fileordirPath string) error {
	initialStat, err := os.Stat(fileordirPath)
	if err != nil {
		return err
	}
	for {
		stat, err := os.Stat(fileordirPath)
		if err != nil {
			return err
		}
		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			killProcess()
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
