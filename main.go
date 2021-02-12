package main

import (
	"fmt"
	"godemon/controllers"
	"godemon/models"
	"os"
)

func main() {
	version := "2.0.6"
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name := controllers.LoadCMD("", "")
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name)
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
				models.TimeLog()
				controllers.BuildMod()
				go controllers.ExecMOD()
			} else if modOrFile == "file" {
				models.TimeLog()
				go controllers.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
