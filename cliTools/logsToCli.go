package cliTools

import (
	"github.com/fatih/color"
	"godemon/models"
	"os"
)

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
