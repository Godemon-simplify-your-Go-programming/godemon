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
				cmd := exec.Command("go", "build", "-o", "app")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				cmd = exec.Command("./app")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				cmd.Process.Kill()
			} else if modOrFile == "file" {
				cmd := exec.Command("go", "run", filepath)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				cmd.Process.Kill()
			}

		}(doneChan)
		<-doneChan
	}
}
