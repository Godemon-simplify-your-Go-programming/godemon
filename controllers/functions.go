package controllers

import (
	"godemon/models"
	"os"
	"os/exec"
	"time"
)

func ExecMOD() {
	cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ExecFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func TimeLog() {
	log := time.Now().Format("2006-01-02, 15:04 \n\n")
	log = `Building project: ` + log + `Program result: `
	cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
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
