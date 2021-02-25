package cliTools

import (
	"flag"
	"github.com/fatih/color"
	"go/build"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, bool, string, string, string, string) {
	var filepathP *string
	var modOrFileP *string
	cnfP := flag.String("cnf", "", "a string")
	filepathP = flag.String("path", "", "a string")
	commandP := flag.String("command", "", "a string")
	helpP := flag.Bool("help", false, "a bool")
	initP := flag.Bool("init", false, "a bool")
	nameP := flag.String("name", "", "a string")
	osP := flag.String("os", "", "a string")
	archP := flag.String("arch", "", "a string")
	modOrFileP = flag.String("modOrFile", "", "a string")
	flag.Parse()
	cnf := *cnfP
	filepath = *filepathP
	init := *initP
	name := *nameP
	os := *osP
	arch := *archP
	command := *commandP
	modOrFile = *modOrFileP
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
	return filepath, modOrFile, cnf, command, helpP, init, name, os, arch, cont
}
