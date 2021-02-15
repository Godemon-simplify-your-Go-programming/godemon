package controllers

import (
	"encoding/json"
	"fmt"
	"godemon/models"
	"io/ioutil"
	"os"
	"time"
)

func ExecMOD() {
	jsonFile, err := os.Open("project.json")
	ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	ErrorHandle(err)
	var pr models.Project
	err = json.Unmarshal(byteValue, &pr)
	ErrorHandle(err)
	for i := 0; i < len(pr.Vars); i++ {
		err = os.Setenv(pr.Vars[i].Key, pr.Vars[i].Value)
		ErrorHandle(err)
	}
	execMOD()
}

func ExecFile(filepath string) {
	execFile(filepath)
}

func TimeLog() {
	timeLog()
}

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, version string, init bool, name string, oso string, arch string) (string, string) {
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile = cnfFunc(command, "", "")
	} else if *cnf == "deploy" {
		deploy(oso, arch)
	} else if init == true {
		if arch == "" && oso == "" {
			fmt.Println("\nPlease specify OS architecture and OS platform")
			os.Exit(1)
		} else if arch == "" {
			fmt.Println("\nPlease specify OS architecture")
			os.Exit(1)
		} else if oso == "" {
			fmt.Println("\nPlease specify OS platform")
			os.Exit(1)
		}
		initialize(name, arch, oso)
	} else if *help == true ||
		(*cnf == "" && filepath == "" && modOrFile == "" &&
			command == "" && *help == false && init == false &&
			name == "" && oso == "" && arch == "") {
		models.HelpCLI(version)
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
