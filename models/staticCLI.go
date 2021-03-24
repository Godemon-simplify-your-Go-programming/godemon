package models

import (
	"fmt"
	"github.com/fatih/color"
	"godemon/godemonInfo"
	"os"
)

func HelpCLI() {
	version := godemonInfo.LoadVersion()
	color.Green("Godemon %v:", version)
	fmt.Printf("\nWhat flags do we have in Godemon? \n	1. Methods of godemon's runtime:" +
		"\n		a) -cnf - using project.json file to load godemon project specifications,\n		b) -cmd - manualy passing godemon project specification using the CLI,\n		c) -deploy - deploying the project using specifications from project.json or CLI\n" +
		"	2. Other flags:\n		a) -path - path to file or dir,\n		b) -mod - specyfing that runtime option is using go modules,\n" +
		"		c) -file - specyfing that runtime option is using single go file,\n		d) -os - specyfing the GOOS,\n" +
		"		e) -arch - specyfing the GOARCH,\n		f) -name - specyfing the name of project/command/file\n" +
		"	3. project.json initialization and update:\n		a) -init - initializating the project.json and go module,\n" +
		"		b) -addCmd - adding command to project.json,\n		c) -addFile - adding file to project.json") //("\n 1. -cnf <- in this flag put info about what do you want to do  - if use cmd just -cmd option, if config file use -cnf \n2. -path <- path to file/directory \n 3. -modOrFile <- are you using modules or one file \n  4. -command <- binded command in config file \n 5. -name <- name of project \n 6. -os <- OS platform \n 7. -arch <- architecture \n 8. -init <- initializazition of project \n\n Have a problem? Go on my GitHub and see the README.md - https://github.com/Godemon-simplify-your-Go-programming/godemon \n")
	os.Exit(1)
}
