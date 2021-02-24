package errors

import (
	"encoding/json"
	"godemon/models"
	"io/ioutil"
	"time"
)

func errorHandle(err error) {
	if err != nil {
		secretErrorHandle(err)
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
