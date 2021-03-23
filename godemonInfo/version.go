package godemonInfo

import (
	"github.com/fatih/color"
	"go/build"
	"os"
)

func LogVersion(version string) {
	fullVersion := "GODEMON-" + version + "-" + build.Default.GOOS + "_" + build.Default.GOARCH
	color.Yellow(fullVersion)
	os.Exit(1)
}
