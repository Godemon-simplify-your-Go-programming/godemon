package godemonInfo

import (
	"os"
	"os/exec"
)

func LogChanges() {

	godemonPath := os.Getenv("GODEMON")
	path := godemonPath + "/CHANGELOGS/Changes.txt"
	cmd := exec.Command("cat", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Exit(1)
}
