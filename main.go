package main

import (
	"fmt"
	"github.com/fatih/color"
	"go/build"
	"godemon/controllers"
	"godemon/models"
	"os"
	"os/exec"
)

func main() {
	hostInfo := [2]string{build.Default.GOOS, build.Default.GOARCH}
	color.Blue("Godemon starting...")
	version := "2.5.4"
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name, oso, arch := controllers.LoadCMD("", "")
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name, oso, arch, hostInfo[0])
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()
			err := controllers.WatchFiles(filepath, hostInfo[0])
			controllers.ErrorHandle(err)
			fmt.Println("File has been changed")
			if modOrFile == "mod" {
				err = os.Chdir(filepath)
				controllers.ErrorHandle(err)
				controllers.TimeLog()
				var cmd *exec.Cmd
				cmd = models.CMDhotReload(hostInfo)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				controllers.ErrorHandle(err)
				go controllers.ExecMOD(hostInfo[0])
			} else if modOrFile == "file" {
				controllers.TimeLog()
				go controllers.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
