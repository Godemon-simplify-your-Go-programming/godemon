package controllers

import (
	"encoding/json"
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
	json.Unmarshal(byteValue, &project)
	return project
}
