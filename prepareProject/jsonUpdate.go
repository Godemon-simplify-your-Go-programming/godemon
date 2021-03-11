package prepareProject

import (
	"encoding/json"
	"github.com/fatih/color"
	"godemon/errors"
	"godemon/models"
	"io/ioutil"
	os2 "os"
)

func ModifyJSONCommands(option string, name string, path string) {
	project := LoadProjectInfo()
	var command models.Command
	command.Name = name
	command.Option = option
	command.Path = path
	project.Commands = append(project.Commands, command)
	file, err := json.MarshalIndent(project, "", "	")
	errors.ErrorHandle(err)
	err = ioutil.WriteFile("project.json", file, 0644)
	errors.ErrorHandle(err)
}

func ModifyJSONFiles(name string, path string) {
	project := LoadProjectInfo()
	var fileJ models.File
	fileJ.Name = name
	fileJ.Path = path
	project.Files = append(project.Files, fileJ)
	file, err := json.MarshalIndent(project, "", "	")
	errors.ErrorHandle(err)
	err = ioutil.WriteFile("project.json", file, 0644)
	errors.ErrorHandle(err)
	color.Green("Everything done, data updated")
	os2.Exit(1)
}

func ModifyJSONVars(key string, value string) {
	project := LoadProjectInfo()
	var fileJ models.Var
	fileJ.Key = key
	fileJ.Value = value
	project.Vars = append(project.Vars, fileJ)
	file, err := json.MarshalIndent(project, "", "	")
	errors.ErrorHandle(err)
	err = ioutil.WriteFile("project.json", file, 0644)
	errors.ErrorHandle(err)
	color.Green("Everything done, data updated")
	os2.Exit(1)
}

func ModifyJSONInfo(name string, os string, arch string, option string) {
	project := LoadProjectInfo()
	if option == "name" {
		project.Name = name
	} else if option == "arch" {
		project.Arch = arch
	} else if os == "os" {
		project.OS = os
	}
	file, err := json.MarshalIndent(project, "", "	")
	errors.ErrorHandle(err)
	err = ioutil.WriteFile("project.json", file, 0644)
	errors.ErrorHandle(err)
	color.Green("Everything done, data updated")
	os2.Exit(1)
}
