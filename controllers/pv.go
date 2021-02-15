package controllers

import (
	"encoding/json"
	"fmt"
	"godemon/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func deploy(oso string, archL string) {
	var goos string
	var arch string
	pr := loadProjectInfo()
	name := pr.Name
	err := os.Chdir(pr.Path)
	ErrorHandle(err)
	if oso != "" && archL == "" {
		goos = "GOOS=" + oso
		arch = "GOARCH=" + archL
	} else {
		goos = "GOOS=" + pr.OS
		arch = "GOARCH=" + pr.Arch
	}
	cmd := exec.Command("env", goos, arch, "go", "build", "-o", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	ErrorHandle(err)
	os.Exit(1)
}

func initialize(name string, arch string, oso string) {
	err := os.Mkdir(name, 0777)
	ErrorHandle(err)
	err = os.Chdir(name)
	ErrorHandle(err)
	var project models.Project
	path, _ := os.Getwd()
	project.Path = path
	project.Name = name
	project.Arch = arch
	project.OS = oso
	project.Vars = append(project.Vars, models.Var{"", ""})
	var command models.Command
	command.Name = "run"
	command.Path = path
	command.Option = "mod"
	project.Commands = append(project.Commands, command)
	file, err := json.MarshalIndent(project, "", "	")
	ErrorHandle(err)
	err = ioutil.WriteFile("project.json", file, 0644)
	ErrorHandle(err)
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	ErrorHandle(err)
	os.Exit(1)
}

func cnfFunc(command string, filepath string, modOrFile string) (string, string) {
	project := loadProjectInfo()
	for i := 0; i < len(project.Commands); i++ {
		if command == project.Commands[i].Name {
			filepath = project.Commands[i].Path
			modOrFile = project.Commands[i].Option
		}
	}
	if filepath == "" || modOrFile == "" {
		fmt.Println("Filepath or modOrFile is empty")
		os.Exit(1)
	}
	return filepath, modOrFile
}
