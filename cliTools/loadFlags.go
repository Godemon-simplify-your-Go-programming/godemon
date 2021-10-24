package cliTools

import (
	"flag"
	"go/build"
	"godemon/godemonInfo"

	"github.com/fatih/color"
)

func LoadCMD(filepath string, modOrFile string) (string, string, string, string, *bool, bool, string, string, string, string, bool, bool, bool, string, string, bool, bool, bool) {
	var filepathP *string
	cmd := flag.Bool("cmd", false, "Manualy passing godemon project specification using the CLI.")
	cnfM := flag.Bool("cnf", false, "Using project.json file to load godemon project specifications.")
	deploy := flag.Bool("deploy", false, "Deploying the project using specifications from project.json or CLI")
	filepathP = flag.String("path", "", "Path to file or dir")
	commandP := flag.String("command", "", "command's name (in project.json)")
	helpP := flag.Bool("help", false, "Help section")
	initP := flag.Bool("init", false, "Initialize project")
	nameP := flag.String("name", "", "Name of project/file/command")
	osP := flag.String("os", "", "OS variable")
	archP := flag.String("arch", "", "Architecture variable")
	mod := flag.Bool("mod", false, "Using as module")
	file := flag.Bool("file", false, "Using as single file")
	addFileM := flag.Bool("addFile", false, "Add file to watching list in project.json")
	addCommandM := flag.Bool("addCommand", false, "Add command to commands list in project.json")
	addVariable := flag.Bool("addVariable", false, "Add tmp variable to project.json")
	value := flag.String("value", "", "value of developer variable")
	key := flag.String("key", "", "key of developer variable")
	updateName := flag.Bool("updateName", false, "Update name of project")
	updateArch := flag.Bool("updateArch", false, "Update arch of project")
	updateOS := flag.Bool("updateOS", false, "Update os of project")
	versionF := flag.Bool("version", false, "Print version of godemon")
	changes := flag.Bool("logChanges", false, "Print changes")
	flag.Parse()
	if *versionF {
		godemonInfo.LogVersion()
	}
	if *changes {
		godemonInfo.LogChanges()
	}
	cnf := ""
	modOrFile = ""
	if *mod {
		modOrFile = "mod"
	} else if *file {
		modOrFile = "mod"
	}
	addFile := false
	if *cmd {
		cnf = "cmd"
	} else if *cnfM {
		cnf = "cnf"
	} else if *deploy {
		cnf = "deploy"
	} else if (*cmd && *cnfM) || (*cmd && *deploy) || (*cnfM && *deploy) {
		color.Red("Too many mode parameters")
	} else if *addFileM {
		addFile = true
	}
	filepath = *filepathP
	init := *initP
	name := *nameP
	os := *osP
	arch := *archP
	command := *commandP

	cont := ""
	if init && name == "" {
		color.Red("Missing parameter: -name")
		cont = "Exit"
	} else if init {
		if arch == "" && os == "" {
			arch = build.Default.GOARCH
			os = build.Default.GOOS
		} else if arch == "" {
			arch = build.Default.GOARCH
		} else if os == "" {
			os = build.Default.GOOS
		}
	}
	if *deploy && (*file || *cmd || *cnfM || *filepathP != "" || *commandP != "" || *helpP || *initP || *nameP != "" || *osP != "" || *archP != "" || *mod) {
		color.Yellow("Warning!!! Deploy takes only deploy argument")
	}
	if *cnfM && *commandP == "" {
		color.Red("You must specify a command")
	}
	return filepath, modOrFile, cnf, command, helpP, init, name, os, arch, cont, addFile, *addCommandM, *addVariable, *key, *value, *updateName, *updateArch, *updateOS
}
