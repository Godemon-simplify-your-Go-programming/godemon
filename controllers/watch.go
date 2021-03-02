package controllers

import (
	"godemon/errors"
	"godemon/killProcess"
	"godemon/prepareProject"
	"os"
	"strings"
	"time"
)

func WatchFiles(fileordirPath string, hOS string) error {
	return watch(fileordirPath, hOS)
}

func watch(fileordirPath string, hOS string) error {
	files := prepareProject.LoadProjectInfo().Files
	j := 0
	var initStats []os.FileInfo
	for j < len(files) {
		stat, err := os.Stat(strings.TrimSpace(files[j].Path))
		errors.ErrorHandle(err)
		initStats = append(initStats, stat)
		j++
	}
	g := 0
	for {
		i := 0
		for i < len(files) {
			stat, err := os.Stat(strings.TrimSpace(files[i].Path))
			if err != nil {
				return err
			}
			if stat.Size() != initStats[i].Size() || stat.ModTime() != initStats[i].ModTime() {
				killProcess.KillProcess(hOS)
				g = 1
				break
			}
			i++
		}
		if g == 1 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
