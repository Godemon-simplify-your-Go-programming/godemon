package models

import (
	"flag"
	"godemon/godemonInfo"
	"os"

	"github.com/fatih/color"
)

func HelpCLI() {
	version := godemonInfo.LoadVersion()
	color.Green("Godemon %v:", version)
	flag.Usage()
	os.Exit(1)
}
