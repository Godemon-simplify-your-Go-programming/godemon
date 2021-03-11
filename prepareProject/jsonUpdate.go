package prepareProject

import (
	"encoding/json"
	"godemon/errors"
	"godemon/models"
	"io/ioutil"
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
}
