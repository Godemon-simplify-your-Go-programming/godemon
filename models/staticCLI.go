package models

import (
	"fmt"
	"os"
)

func HelpCLI(version string) {
	fmt.Printf("Godemon %v: \n 1. -cnf <- in this flag put info about what do you want to do - if use cmd option use -cnf=cmd, if config file use -cnf=cnf \n 2. -path <- path to file/directory \n 3. -modOrFile <- are you using modules or one file \n 4. -command <- binded command in config file \n", version)
	os.Exit(1)
}
