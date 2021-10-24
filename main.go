//TODO - create Windows installer
//TODO - create new README.md
//TODO - create new WIKI

//TODO - update installer - add changelogs directory
//TODO - update updater - add changelogs updating
//TODO - creating version.txt in .godemon/.infos

//TODO - WebPage

//TODO - 21.12 - user errors
//TODO - 21.12 - windows installer

//TODO - 21.12 - create error messages

//!!! - REAPAIR WINDOWS VARIABLES - GODEMON and PATH
package main

import (
	"fmt"
	"go/build"
	"godemon/cliTools"
	"godemon/controllers"
	"godemon/errors"
	"godemon/execs"
	"godemon/godemonInfo"
	"godemon/hotReload"
	"godemon/updateInfo"
	"os"

	"github.com/fatih/color"
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
	filepath, modOrFile, flagsC := controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, init, name, oso, arch, hostInfo[0], addFile)
	color.Cyan("Godemon starting...")
	color.HiMagenta("Welcome to godemon " + version)
	for {
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
				cmd := hotReload.CMDhotReload(hostInfo)
				err = cmd.Run()
				if !cmd.ProcessState.Success() {
					color.Red("\nResult of last working version of project: \n")
				}
				errors.ErrorHandle(err)
				go execs.ExecMOD(hostInfo[0], flagsC)
			} else if modOrFile == "file" {
				cliTools.TimeLog()
				go execs.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
