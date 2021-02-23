package controllers

import (
	"godemon/killProcess"
	"os"
	"time"
)

func WatchFiles(fileordirPath string, hOS string, name string) error {
	return watch(fileordirPath, hOS, name)
}

func watch(fileordirPath string, hOS string, name string) error {
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
			killProcess.KillProcess(hOS, name)
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
