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

func ProgramStarting(cnf string, filepath string, modOrFile string) (string, string) {
	if cnf == "cmd" {
		filepath = os.Args[2]
		modOrFile = os.Args[3]
	} else if cnf == "cnf" {
		command := os.Args[2]
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
