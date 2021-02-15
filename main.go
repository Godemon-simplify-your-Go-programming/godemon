package main

import (
	"fmt"
	"github.com/fatih/color"
	"godemon/controllers"
	"os"
	"os/exec"
)

func main() {
	//hostInfo := [2]string{build.Default.GOOS, build.Default.GOARCH}
	color.Blue("Godemon starting...")
	version := "2.2.2"
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name, oso, arch := controllers.LoadCMD("", "")
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name, oso, arch)
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()
			err := controllers.WatchFiles(filepath)
			controllers.ErrorHandle(err)
			fmt.Println("File has been changed")
			if modOrFile == "mod" {
				err = os.Chdir(filepath)
				controllers.ErrorHandle(err)
				controllers.TimeLog()
				cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				controllers.ErrorHandle(err)
				go controllers.ExecMOD()
			} else if modOrFile == "file" {
				controllers.TimeLog()
				go controllers.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
