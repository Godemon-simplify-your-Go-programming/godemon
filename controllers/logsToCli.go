package controllers

import (
	"os"
	"os/exec"
	"time"
)

func timeLog() {
	log := time.Now().Format("2006-01-02, 15:04 \n\n")
	log = `Building project: ` + log + `Program result: `
	cmd := exec.Command("printf", "\\e[1;34m%-6s\\e[m\n", log)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}
