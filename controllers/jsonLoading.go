package controllers

import (
	"encoding/json"
	"fmt"
	"go/build"
	"godemon/models"
	"io/ioutil"
	"os"
)

func loadProjectInfo() models.Project {
	jsonFile, err := os.Open("project.json")
	ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	ErrorHandle(err)
	var project models.Project
	err = json.Unmarshal(byteValue, &project)
	ErrorHandle(err)
	if project.Name == "" || project.Path == "" {
		fmt.Println("Project name or path is empty")
		os.Exit(1)
	}
	if project.OS == "" && project.Arch == "" {
		project.OS = build.Default.GOOS
		project.Arch = build.Default.GOARCH
	} else if project.OS == "" {
		project.OS = build.Default.GOOS
	} else if project.Arch == "" {
		project.Arch = build.Default.GOARCH
	}
	return project
}
