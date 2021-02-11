package main

import (
	"encoding/json"
	"fmt"
	"godemon/controllers"
	"godemon/models"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func watch(fileordirPath string) error {
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

func main() {
	doneChan := make(chan bool)
	cnf := os.Args[1]
	var filepath string
	var modOrFile string
	if cnf == "cmd" {
		filepath = os.Args[2]
		modOrFile = os.Args[3]
	} else if cnf == "cnf" {
		command := os.Args[2]
		jsonFile, _ := os.Open("godemon-cnf.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)
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
		jsonFile.Close()
	}
	fmt.Println(filepath)
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()

			err := watch(filepath)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("File has been changed")
			if modOrFile == "mod" {
				os.Chdir(filepath)
				controllers.TimeLog()
				cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				go controllers.ExecMOD()
			} else if modOrFile == "file" {
				controllers.TimeLog()
				go controllers.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
