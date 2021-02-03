package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func watchFile(filePath string) error {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
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
	for true {
		go func(doneChan chan bool) {
			defer func() {
				doneChan <- true
			}()

			err := watchFile("./app/main.go")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("File has been changed")
			cmd := exec.Command("go", "run", "./app/main.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}(doneChan)

		<-doneChan
	}
}
