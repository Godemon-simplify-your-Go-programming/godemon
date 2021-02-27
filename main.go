//TODO - optimize memory and cpu/gpu usage
//TODO - add windows full support
//TODO - project.json - can't run single file cuz there's no path variable
//TODO - update as a bin file - godemon-update
//TODO - change CLI logo in installer to lighting
//TODO - create Windows installer
//TODO - auto add of export statement to .zshenv or .bashrc
//TODO - auto add of PATH to PATH in windows
//TODO - golang delve support

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
	"os"
	"os/exec"
)

func main() {
	hostInfo := [2]string{build.Default.GOOS, build.Default.GOARCH}
	color.Cyan("Godemon starting...")
	version := "21.04"
	color.HiMagenta("Welcome to godemon " + version)
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name, oso, arch, cont := cliTools.LoadCMD("", "")
	if cont == "Exit" {
		os.Exit(1)
	}
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name, oso, arch, hostInfo[0])
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
