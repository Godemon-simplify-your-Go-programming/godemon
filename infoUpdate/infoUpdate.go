package infoUpdate

import (
	"fmt"
	"godemon/prepareProject"
	"os"
)

func Update(addCmd bool, addFile bool, name string, modOrFile string, filepath string) {
	if addCmd == true && name != "" && modOrFile != "" {
		prepareProject.ModifyJSONCommands(modOrFile, name, filepath)
		fmt.Println("Info updated")
		os.Exit(1)
	}
	if addFile == true && name != "" && filepath != "" {
		prepareProject.ModifyJSONFiles(name, filepath)
		fmt.Println("Info updated")
		os.Exit(1)
	}
}
