package controllers

import (
	"encoding/json"
	"fmt"
	"godemon/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func killProcess() {
	cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func cnfCnfLoad(projectInfo models.ProjectInfo, command string, filepath string, modOrFile string) (string, string) {
	commandsL := projectInfo.Commands.Commands
	for i := 0; i < len(commandsL); i++ {
		if command == commandsL[i].Name {
			fmt.Println(commandsL[i].Path)
			fmt.Println(commandsL[i].Option)
			filepath = commandsL[i].Path
			modOrFile = commandsL[i].Option
		}
	}
	return filepath, modOrFile
}

func cnfDepLoad(projectInfo models.ProjectInfo, command string, filepath string, modOrFile string) {
	models.DeployLog()
	path := projectInfo.Project.Path
	osP := projectInfo.Project.OS
	arch := projectInfo.Project.Arch
	goos := "GOOS=" + osP
	archos := "GOARCH=" + arch
	name := projectInfo.Project.Name
	os.Chdir(path)
	cmd := exec.Command("env", goos, archos, "go", "build", "-o", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", "Project builded")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Exit(1)
}

func loadDataFromCnf(cnf *string, command string, filepath string, modOrFile string) (string, string) {
	jsonFile, err := os.Open("godemon-cnf.json")
	ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	ErrorHandle(err)
	var projectInfo models.ProjectInfo
	json.Unmarshal(byteValue, &projectInfo)
	if *cnf == "cnf" {
		cnfCnfLoad(projectInfo, command, "", "")
	} else if *cnf == "deploy" {
		cnfDepLoad(projectInfo, command, "", "")
	}
	return filepath, modOrFile
}

func initialization(name string) {
	os.Mkdir(name, 0777)
	os.Chmod(name, 0777)
	path, _ := os.Getwd()
	dir := path + "/" + name
	fmt.Println(dir)
	os.Chdir(dir)
	path, _ = os.Getwd()
	var p models.Project
	p.Path = path
	p.Name = name
	var c models.Commands
	projectInfo := models.ProjectInfo{p, c}
	file, _ := json.MarshalIndent(projectInfo, "", "")
	f := dir + "/godemon-cnf.json"
	fmt.Println(f)
	_ = ioutil.WriteFile(f, file, 0777)
	models.InitLog()
	os.Exit(1)
}
