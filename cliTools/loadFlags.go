package cliTools

import (
	"flag"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, bool, string, string, string, bool) {
	var filepathP *string
	var modOrFileP *string
	cnfP := flag.String("cnf", "", "a string")
	filepathP = flag.String("path", "", "a string")
	commandP := flag.String("command", "", "a string")
	helpP := flag.Bool("help", false, "a bool")
	logsP := flag.Bool("logs", false, "a bool")
	initP := flag.Bool("init", false, "a bool")
	nameP := flag.String("name", "", "a string")
	osP := flag.String("os", "", "a string")
	archP := flag.String("arch", "", "a string")
	modOrFileP = flag.String("modOrFile", "", "a string")
	flag.Parse()
	cnf := *cnfP
	logs := *logsP
	filepath = *filepathP
	init := *initP
	name := *nameP
	os := *osP
	arch := *archP
	command := *commandP
	modOrFile = *modOrFileP
	return filepath, modOrFile, cnf, command, helpP, init, name, os, arch, logs
}
