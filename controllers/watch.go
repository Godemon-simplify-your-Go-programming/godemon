package controllers

import (
	"godemon/killProcess"
	"os"
	"time"
)

func WatchFiles(fileordirPath string, hOS string) error {
	return watch(fileordirPath, hOS)
}

func watch(fileordirPath string, hOS string) error {
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
			killProcess.KillProcess(hOS)
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
