package execs

import (
	"encoding/json"
	"godemon/errors"
	"godemon/models"
	"io/ioutil"
	"os"
)

func ExecMOD(hOS string, name string, logs bool) {
	jsonFile, err := os.Open("project.json")
	errors.ErrorHandle(err)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	errors.ErrorHandle(err)
	var pr models.Project
	err = json.Unmarshal(byteValue, &pr)
	errors.ErrorHandle(err)
	for i := 0; i < len(pr.Vars); i++ {
		err = os.Setenv(pr.Vars[i].Key, pr.Vars[i].Value)
		errors.TMPerrorHandle(err)
	}
	execMOD(hOS, name, logs)
}

func ExecFile(filepath string) {
	execFile(filepath)
}
