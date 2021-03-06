//TODO - create Windows installer - Patryk
//TODO - create new README.md - Norbert
//TODO - create new WIKI - Norbert
//TODO - flags in commands

//TODO - WebPage

//TODO - 21.06 - golang delve support
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
	"godemon/models"
	"godemon/updateInfo"
	"os"
	"os/exec"
)

func main() {
	hostInfo := [2]string{build.Default.GOOS, build.Default.GOARCH}
	color.Cyan("Godemon starting...")
	version := "21.06"
	color.HiMagenta("Welcome to godemon " + version)
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name, oso, arch, cont, addFile, addCmd, addVar, key, value, updateName, updateArch, updateOS := cliTools.LoadCMD("", "")
	if cont == "Exit" {
		os.Exit(1)
	}
	updateInfo.Update(updateName, name, updateArch, arch, oso, updateOS, addVar, key, value, addCmd, modOrFile, filepath)
	var flags []models.Flag
	filepath, modOrFile, flags = controllers.ProgramStarting(flags, &cnf, filepath, modOrFile, command, help, version, init, name, oso, arch, hostInfo[0], addFile)
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()
			err := controllers.WatchFiles(filepath, hostInfo[0])
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
				go execs.ExecMOD(hostInfo[0], flags)
			} else if modOrFile == "file" {
				cliTools.TimeLog()
				go execs.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
