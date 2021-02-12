package main

import (
	"fmt"
	"godemon/controllers"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", "Starting godemon... \n")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	version := "2.1.2"
	doneChan := make(chan bool)
	filepath, modOrFile, cnf, command, help, init, name, oso, arch := controllers.LoadCMD("", "")
	filepath, modOrFile = controllers.ProgramStarting(&cnf, filepath, modOrFile, command, help, version, init, name, oso, arch)
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
				os.Chdir("/home/nwagner/Desktop/godemon/godemon_test/")
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
