package main

import (
	"fmt"
	"godemon/controllers"
	"os"
	"os/exec"
)

func main() {
	version := "2.0.0"
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help := controllers.LoadCMD("", "")
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version)
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()
			err := controllers.WatchFiles(filepath)
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
