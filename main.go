//TODO - create Windows installer - Patryk
//TODO - create new README.md - Norbert
//TODO - create new WIKI - Norbert
//TODO - flags in commands
//TODO - update installer - add changelogs directory
//TODO - update updater - add changelogs updating
//TODO - creating version.txt in .godemon/.infos

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
	"godemon/godemonInfo"
	"godemon/hotReload"
	"godemon/updateInfo"
	"os"
	"os/exec"
)

func main() {
	hostInfo := [2]string{build.Default.GOOS, build.Default.GOARCH}
	doneChan := make(chan bool)
	version := godemonInfo.LoadVersion()
	filepath, modOrFile, cnf, command, help, init, name, oso, arch, cont, addFile, addCmd, addVar, key, value, updateName, updateArch, updateOS := cliTools.LoadCMD("", "")
	if cont == "Exit" {
		os.Exit(1)
	}
	updateInfo.Update(updateName, name, updateArch, arch, oso, updateOS, addVar, key, value, addCmd, modOrFile, filepath)
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, init, name, oso, arch, hostInfo[0], addFile)
	color.Cyan("Godemon starting...")
	color.HiMagenta("Welcome to godemon " + version)
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
