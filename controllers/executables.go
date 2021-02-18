package controllers

import (
	"encoding/json"
	"godemon/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func ExecMOD(hOS string) {
	jsonFile, err := os.Open("project.json")
	ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	ErrorHandle(err)
	var pr models.Project
	err = json.Unmarshal(byteValue, &pr)
	ErrorHandle(err)
	for i := 0; i < len(pr.Vars); i++ {
		err = os.Setenv(pr.Vars[i].Key, pr.Vars[i].Value)
		ErrorHandle(err)
	}
	execMOD(hOS)
}

func ExecFile(filepath string) {
	execFile(filepath)
}

func execMOD(hOS string) {
	if hOS == "windows" {
		cmd := exec.Command("app-godemon-app-godemon-tmp-generated.exe")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	} else {
		cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	}
}

func execFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}
