package controllers

import (
	"flag"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, *bool, string) {
	var filepathP *string
	var modOrFileP *string
	cnfP := flag.String("cnf", "", "a string")
	filepathP = flag.String("path", "", "a string")
	commandP := flag.String("command", "", "a string")
	helpP := flag.Bool("help", false, "a bool")
	initP := flag.Bool("init", false, "a bool")
	nameP := flag.String("name", "", "a string")
	modOrFileP = flag.String("modOrFile", "", "a string")
	flag.Parse()
	cnf := *cnfP
	filepath = *filepathP
	command := *commandP
	name := *nameP
	modOrFile = *modOrFileP
	if cnf == "" {
		*helpP = true
	}
	return filepath, modOrFile, cnf, command, helpP, initP, name
}
