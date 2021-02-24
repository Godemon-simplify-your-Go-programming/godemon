package errors

import (
	"encoding/json"
	"github.com/fatih/color"
	"godemon/models"
	"io/ioutil"
	"time"
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
		fileName := "error" + time.Now().Format("2006-01-02") + ".json"
		err = ioutil.WriteFile("~/.godemon/logs/"+fileName, file, 0644)
	}
}
