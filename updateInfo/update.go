package updateInfo

import (
	"github.com/fatih/color"
	"godemon/prepareProject"
	"os"
)

func Update(updateName bool, name string, updateArch bool, arch string, oso string, updateOS bool, addVar bool, key string, value string, addCmd bool, modOrFile string, filepath string) {
	option := ""
	if updateName == true {
		if name == "" {
			color.Red("Name is empty")
			os.Exit(1)
		}
		option = "name"
		prepareProject.ModifyJSONInfo(name, oso, arch, option)
	} else if updateArch == true {
		if arch == "" {
			color.Red("Arch parameter is empty")
			os.Exit(1)
		}
		option = "arch"
		prepareProject.ModifyJSONInfo(name, oso, arch, option)
	} else if updateOS == true {
		if oso == "" {
			color.Red("OS parameter is empty")
			os.Exit(1)
		}
		option = "os"
		prepareProject.ModifyJSONInfo(name, oso, arch, option)
	}
	if addVar == true {
		if key == "" || value == "" {
			color.Red("Key or value is empty")
			os.Exit(1)
		}
		prepareProject.ModifyJSONVars(key, value)
	}
	if addCmd == true {
		if name == "" || modOrFile == "" {
			color.Red("Name is empty or option isn't specified")
		}
		prepareProject.ModifyJSONCommands(modOrFile, name, filepath)
	}
}
