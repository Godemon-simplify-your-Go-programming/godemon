package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func watch(fileordirPath string) error {
	initialStat, err := os.Stat(fileordirPath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(fileordirPath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func execMOD() {
	cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func execFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func timeLog() {
	log := time.Now().Format("2006-01-02, 15:04 \n\n")
	log = `Building project: ` + log + `Program result: `
	cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	doneChan := make(chan bool)
	filepath := os.Args[1]
	modOrFile := os.Args[2]
	fmt.Println(filepath)
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()

			err := watch(filepath)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("File has been changed")
			if modOrFile == "mod" {
				os.Chdir(filepath)
				timeLog()
				cmd := exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				go execMOD()
			} else if modOrFile == "file" {
				timeLog()
				go execFile(filepath)
			}
		}(doneChan)
		<-doneChan
	}
}
