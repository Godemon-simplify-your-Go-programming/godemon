package controllers

import (
	"fmt"
	"godemon/models"
	"os"
)

//TODO("CREATE NEW MODULES - TimeLog, Executable, JsonLoad, Kill, LogsToCli, LoadFlags - everything with private functions")

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, version string, init bool, name string, oso string, arch string, hOS string) (string, string) {
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile = cnfFunc(command, "", "")
	} else if *cnf == "deploy" {
		deploy(oso, arch, hOS)
	} else if init == true {
		if arch == "" && oso == "" {
			fmt.Println("\nPlease specify OS architecture and OS platform")
			os.Exit(1)
		} else if arch == "" {
			fmt.Println("\nPlease specify OS architecture")
			os.Exit(1)
		} else if oso == "" {
			fmt.Println("\nPlease specify OS platform")
			os.Exit(1)
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
