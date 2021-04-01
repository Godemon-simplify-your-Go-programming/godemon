package controllers

import (
	"github.com/fatih/color"
	"godemon/models"
	"os"
)

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, init bool, name string, oso string, arch string, hOS string, addFile bool) (string, string, []string) {
	var flagsC []string
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile, flagsC = cnfFunc(command, "", "")
	} else if *cnf == "deploy" {
		deploy(oso, arch, hOS)
	} else if init == true {
		initialize(name, arch, oso)
	} else if addFile == true {
		if name == "" || filepath == "" {
			color.Red("Name or filepath is empty")
			os.Exit(1)
		}
		addFileJson(name, filepath)
	} else if *help == true ||
		(*cnf == "" && filepath == "" && modOrFile == "" &&
			command == "" && *help == false && init == false &&
			name == "" && oso == "" && arch == "") {
		models.HelpCLI()
	}
	return filepath, modOrFile, flagsC
}
