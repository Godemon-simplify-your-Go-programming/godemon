//TODO - create Windows installer - Patryk
//TODO - create new README.md - Norbert
//TODO - create new WIKI - Norbert
//TODO - new help section - Patryk

//TODO - 21.05 - golang delve support
package main

import (
	"fmt"
	"github.com/fatih/color"
	"go/build"
	"godemon/cliTools"
	"godemon/controllers"
	"godemon/errors"
	"godemon/execs"
	"godemon/hotReload"
	"godemon/infoUpdate"
	"os"
	"os/exec"
)

func main() {
	hostInfo := [2]string{build.Default.GOOS, build.Default.GOARCH}
	color.Cyan("Godemon starting...")
	version := "21.04"
	color.HiMagenta("Welcome to godemon " + version)
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name, oso, arch, cont, addCmd, addFile := cliTools.LoadCMD("", "")
	if cont == "Exit" {
		os.Exit(1)
	}
	infoUpdate.Update(addCmd, addFile, name, modOrFile, filepath)
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name, oso, arch, hostInfo[0])
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()
			err := controllers.WatchFiles(filepath, hostInfo[0], cnf)
			errors.ErrorHandle(err)
			fmt.Println("File has been changed")
			if modOrFile == "mod" {
				err = os.Chdir(filepath)
				errors.ErrorHandle(err)
				cliTools.TimeLog()
				var cmd *exec.Cmd
				cmd = hotReload.CMDhotReload(hostInfo)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				errors.ErrorHandle(err)
				go execs.ExecMOD(hostInfo[0])
			} else if modOrFile == "file" {
				cliTools.TimeLog()
				go execs.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
