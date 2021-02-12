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

func ExecMOD() {
	cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ExecFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func TimeLog() {
	log := time.Now().Format("2006-01-02, 15:04 \n\n")
	log = `Building project: ` + log + `Program result: `
	cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, version string) (string, string) {
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile = loadDataFromCnf(cnf, command, "", "")
	} else if *cnf == "deploy" {
		_, _ = loadDataFromCnf(cnf, command, "", "")
	} else if *help == true {
		fmt.Printf("Godemon %v: \n 1. -cnf <- in this flag put info about what do you want to do - if use cmd option use -cnf=cmd, if config file use -cnf=cnf \n 2. -path <- path to file/directory \n 3. -modOrFile <- are you using modules or one file \n 4. -command <- binded command in config file \n", version)
		os.Exit(1)
	}
	return filepath, modOrFile
}

func WatchFiles(fileordirPath string) error {
	initialStat, err := os.Stat(fileordirPath)
	if err != nil {
		return err
	}
	for {
		stat, err := os.Stat(fileordirPath)
		if err != nil {
			return err
		}
		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			killProcess()
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func BuildMod() {
	cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func killProcess() {
	cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
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
		commandsL := projectInfo.Commands.Commands
		for i := 0; i < len(commandsL); i++ {
			if command == commandsL[i].Name {
				fmt.Println(commandsL[i].Path)
				fmt.Println(commandsL[i].Option)
				filepath = commandsL[i].Path
				modOrFile = commandsL[i].Option
			}
		}
	} else if *cnf == "deploy" {
		log := time.Now().Format("2006-01-02, 15:04 \n\n")
		log = `Building project: ` + log + `: `
		cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		path := projectInfo.Project.Path
		osP := projectInfo.Project.OS
		arch := projectInfo.Project.Arch
		goos := "GOOS=" + osP
		archos := "GOARCH=" + arch
		name := projectInfo.Project.Name
		os.Chdir(path)
		cmd = exec.Command("env", goos, archos, "go", "build", "-o", name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		cmd = exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", "Project builded")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(1)
	}
	return filepath, modOrFile
}
