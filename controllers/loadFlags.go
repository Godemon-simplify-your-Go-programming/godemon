package controllers

import (
	"flag"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool) {
	var filepathP *string
	var modOrFileP *string
	cnfP := flag.String("cnf", "", "a string")
	filepathP = flag.String("path", "", "a string")
	commandP := flag.String("command", "", "a string")
	helpP := flag.Bool("help", false, "a bool")
	modOrFileP = flag.String("modOrFile", "", "a string")
	flag.Parse()
	cnf := *cnfP
	filepath = *filepathP
	command := *commandP
	modOrFile = *modOrFileP
	return filepath, modOrFile, cnf, command, helpP
}
