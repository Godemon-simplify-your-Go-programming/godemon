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
	} else if *cnf == "deploy" {
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
		cmd := exec.Command("env", goos, arch, "go", "build", "-o", name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(1)
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
			cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
