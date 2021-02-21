package errors

import (
	"encoding/json"
	"github.com/fatih/color"
	"godemon/models"
	"io/ioutil"
)

func errorHandle(err error) {
	if err != nil {
		color.Cyan(err.Error())
	}
}

func secretErrorHandle(err error) {
	if err != nil {
		var errorM models.ErrorTMP
		file, _ := json.MarshalIndent(errorM, "", "	")
		err = ioutil.WriteFile("/tmp/errorM.json", file, 0644)
	}
}
