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
