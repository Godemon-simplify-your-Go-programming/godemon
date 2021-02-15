package controllers

import (
	"github.com/fatih/color"
	"time"
)

func timeLog() {
	log := time.Now().Format("2006-01-02, 15:04 \n\n")
	log = `Building project: ` + log + `Program result: `
	color.Blue(log)
}
