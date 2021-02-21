package controllers

import (
	"go/build"
	"godemon/models"
)

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, version string, init bool, name string, oso string, arch string, hOS string) (string, string) {
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile = cnfFunc(command, "", "")
	} else if *cnf == "deploy" {
		deploy(oso, arch, hOS)
	} else if init == true {
		if arch == "" && oso == "" {
			arch = build.Default.GOARCH
			oso = build.Default.GOOS
		} else if arch == "" {
			arch = build.Default.GOARCH
		} else if oso == "" {
			oso = build.Default.GOOS
		}
		initialize(name, arch, oso)
	} else if *help == true ||
		(*cnf == "" && filepath == "" && modOrFile == "" &&
			command == "" && *help == false && init == false &&
			name == "" && oso == "" && arch == "") {
		models.HelpCLI(version)
	}
	return filepath, modOrFile
}
