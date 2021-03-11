//TODO - create Windows installer - Patryk
//TODO - create new README.md - Norbert
//TODO - create new WIKI - Norbert

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
	"godemon/prepareProject"
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
	option := ""
	if updateName == true {
		if name == "" {
			color.Red("Name is empty")
			os.Exit(1)
		}
		option = "name"
		prepareProject.ModifyJSONInfo(name, oso, arch, option)
	} else if updateArch == true {
		if arch == "" {
			color.Red("Arch parameter is empty")
			os.Exit(1)
		}
		option = "arch"
		prepareProject.ModifyJSONInfo(name, oso, arch, option)
	} else if updateOS == true {
		if oso == "" {
			color.Red("OS parameter is empty")
			os.Exit(1)
		}
		option = "os"
		prepareProject.ModifyJSONInfo(name, oso, arch, option)
	}
	if addVar == true {
		if key == "" || value == "" {
			color.Red("Key or value is empty")
			os.Exit(1)
		}
		prepareProject.ModifyJSONVars(key, value)
		os.Exit(1)
	}
	if addCmd == true {
		if name == "" || modOrFile == "" {
			color.Red("Name is empty or option isn't specified")
		}
		prepareProject.ModifyJSONCommands(modOrFile, name, filepath)
		os.Exit(1)
	}
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name, oso, arch, hostInfo[0], addFile)
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
				go execs.ExecMOD(hostInfo[0])
			} else if modOrFile == "file" {
				cliTools.TimeLog()
				go execs.ExecFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
