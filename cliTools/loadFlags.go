package cliTools

import (
	"flag"
	"github.com/fatih/color"
	"go/build"
	"godemon/godemonInfo"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, bool, string, string, string, string, bool, bool, bool, string, string, bool, bool, bool) {
	var filepathP *string
	cmd := flag.Bool("cmd", false, "CMD mode")
	cnfM := flag.Bool("cnf", false, "project.json mode")
	deploy := flag.Bool("deploy", false, "building project using info from project.json")
	filepathP = flag.String("path", "", "path to file or dir")
	commandP := flag.String("command", "", "command name")
	helpP := flag.Bool("help", false, "help section")
	initP := flag.Bool("init", false, "initialize project")
	nameP := flag.String("name", "", "name of project/file/command")
	osP := flag.String("os", "", "os variable")
	archP := flag.String("arch", "", "architecture variable")
	mod := flag.Bool("mod", false, "module option")
	file := flag.Bool("file", false, "file option")
	addFileM := flag.Bool("addFile", false, "add file to watching list in project.json")
	addCommandM := flag.Bool("addCommand", false, "add command to commands list in project.json")
	addVariable := flag.Bool("addVariable", false, "add tmp variable to project.json")
	value := flag.String("value", "", "value of ...")
	key := flag.String("key", "", "key of ...")
	updateName := flag.Bool("updateName", false, "update name of project")
	updateArch := flag.Bool("updateArch", false, "update arch of project")
	updateOS := flag.Bool("updateOS", false, "update os of project")
	versionF := flag.Bool("version", false, "print version of godemon")
	changes := flag.Bool("logChanges", false, "print changes")
	flag.Parse()
	if *versionF == true {
		godemonInfo.LogVersion()
	}
	if *changes == true {
		godemonInfo.LogChanges()
	}
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
	return filepath, modOrFile, cnf, command, helpP, init, name, os, arch, cont, addFile, *addCommandM, *addVariable, *key, *value, *updateName, *updateArch, *updateOS
}
