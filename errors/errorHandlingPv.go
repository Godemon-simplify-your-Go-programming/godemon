package errors

import (
	"github.com/fatih/color"
)

func errorHandle(err error) {
	if err != nil {
		color.Cyan(err.Error())
	}
}
