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
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	doneChan := make(chan bool)
	filepath := os.Args[1]
	modOrFile := os.Args[2]
	fmt.Println(filepath)
	var log string
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
			os.Chdir(filepath)

			if modOrFile == "mod" {
				log = time.Now().Format("2006-01-02, 15:04 \n\n")
				log = `Building project: ` + log + `Program result: `
				cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				cmd = exec.Command("go", "build", "-o", "app-godemon-app-godemon-tmp-generated")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				cmd = exec.Command("./app-godemon-app-godemon-tmp-generated")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				cmd = exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			} else if modOrFile == "file" {
				log = time.Now().Format("2006-01-02, 15:04 \n\n")
				log = `Building project: ` + log + `Program result: `
				cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()

				cmd.Process.Kill()
			}

		}(doneChan)
		<-doneChan
	}
}
