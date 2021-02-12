package controllers

import (
	"encoding/json"
	"fmt"
	"godemon/models"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func deploy() {
	jsonFile, err := os.Open("project.json")
	ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	ErrorHandle(err)
	var pr models.Project
	json.Unmarshal(byteValue, &pr)
	goos := "GOOS=" + pr.OS
	arch := "GOARCH=" + pr.Arch
	name := pr.Name
	os.Chdir(pr.Path)
	fmt.Println(pr)
	cmd := exec.Command("env", goos, arch, "go", "build", "-o", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Exit(1)
}

func initialize(name string, arch string, oso string) {
	os.Mkdir(name, 0777)
	os.Chdir(name)
	var project models.Project
	path, _ := os.Getwd()
	project.Path = path
	project.Name = name
	project.Arch = arch
	project.OS = oso
	var commands models.Commands
	var command models.Command
	command.Name = "run"
	command.Path = path
	command.Option = "mod"
	commands.Commands = append(commands.Commands, command)
	file, _ := json.MarshalIndent(project, "", "	")
	_ = ioutil.WriteFile("project.json", file, 0644)
	file, _ = json.MarshalIndent(commands, "", "	")
	_ = ioutil.WriteFile("godemon-cnf.json", file, 0644)
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Exit(1)
}

func cnfFunc(command string, filepath string, modOrFile string) (string, string) {
	jsonFile, err := os.Open("godemon-cnf.json")
	ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	ErrorHandle(err)
	var commands models.Commands
	json.Unmarshal(byteValue, &commands)
	for i := 0; i < len(commands.Commands); i++ {
		if command == commands.Commands[i].Name {
			fmt.Println(commands.Commands[i].Path)
			fmt.Println(commands.Commands[i].Option)
			filepath = commands.Commands[i].Path
			modOrFile = commands.Commands[i].Option
		}
	}
	fmt.Println(filepath, modOrFile)
	time.Sleep(10 * time.Second)
	return filepath, modOrFile
}

func killProcess() {
	cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
