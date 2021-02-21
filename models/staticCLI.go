package models

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func HelpCLI(version string) {
	color.Green("Godemon %v:", version)
	fmt.Println("\n 1. -cnf <- in this flag put info about what do you want to do - if use cmd option use -cnf=cmd, if config file use -cnf=cnf \n 2. -path <- path to file/directory \n 3. -modOrFile <- are you using modules or one file \n 4. -command <- binded command in config file \n 5. -name <- name of project \n 6. -os <- OS platform \n 7. -arch <- architecture \n 8. -init <- initializazition of project \n\n Have a problem? Go on my GitHub and see the README.md - https://github.com/nProgrammer/godemon \n")
	os.Exit(1)
}
