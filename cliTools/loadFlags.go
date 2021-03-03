package cliTools

import (
	"flag"
	"github.com/fatih/color"
	"go/build"
	os1 "os"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, bool, string, string, string, string) {
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
	flag.Parse()
	cnf := ""
	if *cmd == true {
		cnf = "cmd"
	} else if *cnfM == true {
		cnf = "cnf"
	} else if *deploy == true {
		cnf = "deploy"
	} else if (*cmd == true && *cnfM == true) || (*cmd == true && *deploy == true) || (*cnfM == true && *deploy == true) {
		color.Red("Too many mode parameters")
		os1.Exit(1)
	}
	filepath = *filepathP
	init := *initP
	name := *nameP
	os := *osP
	arch := *archP
	command := *commandP
	modOrFile = ""
	if *mod == true {
		modOrFile = "mod"
	} else if *file == true {
		modOrFile = "file"
	}
	if *cmd == true && filepath == "" {
		color.Red("Path is empty")
		os1.Exit(1)
	} else if *cmd == true && *file == false && *mod == false {
		color.Red("You didn't define is it file or module")
		os1.Exit(1)
	}
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
		color.Red("Warning!!! Deploy takes only deploy argument")
		os1.Exit(1)
	}
	return filepath, modOrFile, cnf, command, helpP, init, name, os, arch, cont
}
