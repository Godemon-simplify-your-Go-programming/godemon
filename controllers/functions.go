package controllers

import (
	"godemon/models"
	"os"
	"time"
)

func ProgramStarting(cnf *string, filepath string, modOrFile string, command string, help *bool, version string, init *bool, name string) (string, string) {
	if *cnf == "cmd" {

	} else if *cnf == "cnf" {
		filepath, modOrFile = loadDataFromCnf(cnf, command, "", "")
	} else if *cnf == "deploy" {
		_, _ = loadDataFromCnf(cnf, command, "", "")
	} else if *init == true {
		initialization(name)
	} else if *help == true {
		models.Help(version)
	}
	return filepath, modOrFile
}

func WatchFiles(fileordirPath string) error {
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
			killProcess()
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
