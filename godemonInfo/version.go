package godemonInfo

import (
	"github.com/fatih/color"
	"go/build"
	"io/ioutil"
	"os"
	"strings"
)

func LoadVersion() string {
	godemonPath := os.Getenv("GODEMON")
	path := godemonPath + "/.infos/version.txt"
	v, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(v))
}

func LogVersion() {
	version := LoadVersion()
	fullVersion := "GODEMON-" + version + "-" + build.Default.GOOS + "_" + build.Default.GOARCH
	color.Yellow(fullVersion)
	os.Exit(1)
}
