package cliTools

import (
	"flag"
	"github.com/fatih/color"
	"go/build"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, bool, string, string, string, string, bool, bool, bool, string, string) {
	var filepathP *string
	cmd := flag.Bool("cmd", false, "a bool")
	cnfM := flag.Bool("cnf", false, "a bool")
	deploy := flag.Bool("deploy", false, "a bool")
	filepathP = flag.String("path", "", "a string")
	commandP := flag.String("command", "", "a string")
	helpP := flag.Bool("help", false, "a bool")
	initP := flag.Bool("init", false, "a bool")
	nameP := flag.String("name", "", "a string")
	osP := flag.String("os", "", "a string")
	archP := flag.String("arch", "", "a string")
	mod := flag.Bool("mod", false, "a bool")
	file := flag.Bool("file", false, "a bool")
	addFileM := flag.Bool("addFile", false, "a bool")
	addCommandM := flag.Bool("addCommand", false, "a bool")
	addVariable := flag.Bool("addVariable", false, "a bool")
	value := flag.String("value", "", "a string")
	key := flag.String("key", "", "a string")
	flag.Parse()
	cnf := ""
	modOrFile = ""
	if *mod == true {
		modOrFile = "mod"
	} else if *file == true {
		modOrFile = "mod"
	}
	addFile := false
	if *cmd == true {
		cnf = "cmd"
	} else if *cnfM == true {
		cnf = "cnf"
	} else if *deploy == true {
		cnf = "deploy"
	} else if (*cmd == true && *cnfM == true) || (*cmd == true && *deploy == true) || (*cnfM == true && *deploy == true) {
		color.Red("Too many mode parameters")
	} else if *addFileM == true {
		addFile = true
	}
	filepath = *filepathP
	init := *initP
	name := *nameP
	os := *osP
	arch := *archP
	command := *commandP

	cont := ""
	if init == true && name == "" {
		color.Red("Missing parameter: -name")
		cont = "Exit"
	} else if init == true {
		if arch == "" && os == "" {
			arch = build.Default.GOARCH
			os = build.Default.GOOS
		} else if arch == "" {
			arch = build.Default.GOARCH
		} else if os == "" {
			os = build.Default.GOOS
		}
	}
	if *deploy == true && (*file == true || *cmd == true || *cnfM == true || *filepathP != "" || *commandP != "" || *helpP == true || *initP == true || *nameP != "" || *osP != "" || *archP != "" || *mod == true) {
		color.Yellow("Warning!!! Deploy takes only deploy argument")
	}
	if *cnfM == true && *commandP == "" {
		color.Red("You must specify a command")
	}
	return filepath, modOrFile, cnf, command, helpP, init, name, os, arch, cont, addFile, *addCommandM, *addVariable, *key, *value
}
