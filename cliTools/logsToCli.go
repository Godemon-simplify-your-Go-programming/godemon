package cliTools

import (
	"github.com/fatih/color"
	"godemon/models"
	"os"
	"time"
)

func timeLog() {
	log := time.Now().Format("2006-01-02, 15:04 \n\n")
	log = `Building project: ` + log + `Program result: `
	color.Green(log)
}

func CheckModOrPath(mod string, path string) {
	if path == "" || mod == "" {
		color.Red("Filepath or modOrFile is empty")
		os.Exit(1)
	}
}

func LoadedNameIsEmpty(project models.Project) {
	if project.Name == "" {
		color.Red("Project name is empty")
		os.Exit(1)
	}
}
