package controllers

import (
	"encoding/json"
	"godemon/cliTools"
	"godemon/errors"
	"godemon/models"
	"godemon/prepareProject"
	"io/ioutil"
	"os"
	"os/exec"
)

func deploy(oso string, archL string, hOS string) {
	var goos string
	var arch string
	pr := prepareProject.LoadProjectInfo()
	name := pr.Name
	if oso != "" && archL == "" {
		goos = "GOOS=" + oso
		arch = "GOARCH=" + archL
	} else {
		goos = "GOOS=" + pr.OS
		arch = "GOARCH=" + pr.Arch
	}
	if hOS == "windows" {
		name = name + ".exe"
	}
	cmd := exec.Command("env", goos, arch, "go", "build", "-o", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errors.ErrorHandle(err)
	os.Exit(1)
}

func initialize(name string, arch string, oso string) {
	err := os.Mkdir(name, 0777)
	errors.ErrorHandle(err)
	err = os.Chdir(name)
	errors.ErrorHandle(err)
	var project models.Project
	project.Name = name
	project.Arch = arch
	project.OS = oso
	project.Vars = append(project.Vars, models.Var{"", ""})
	var command models.Command
	command.Name = "run"
	command.Option = "mod"
	project.Commands = append(project.Commands, command)
	file, err := json.MarshalIndent(project, "", "	")
	errors.ErrorHandle(err)
	err = ioutil.WriteFile("project.json", file, 0644)
	errors.ErrorHandle(err)
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	errors.ErrorHandle(err)
	os.Exit(1)
}

func cnfFunc(command string, filepath string, modOrFile string) (string, string, []models.Flag) {
	var err error
	var flags []models.Flag
	iterate := 0
	project := prepareProject.LoadProjectInfo()
	for i := 0; i < len(project.Commands); i++ {
		if command == project.Commands[i].Name {
			filepath, err = os.Getwd()
			errors.ErrorHandle(err)
			modOrFile = project.Commands[i].Option
			if modOrFile == "file" {
				filepath = project.Commands[i].Path
			}
			iterate = i
		}
	}
	j := 0
	for j < len(project.Commands[iterate].Flags) {
		flags[j] = project.Commands[iterate].Flags[j]
		j++
	}
	cliTools.CheckModOrPath(modOrFile, filepath)
	return filepath, modOrFile, flags
}

func addFileJson(name string, path string) {
	prepareProject.ModifyJSONFiles(name, path)
	os.Exit(1)
}
