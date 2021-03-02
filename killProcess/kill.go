package killProcess

import (
	"godemon/errors"
	"os"
)

func KillProcess(hOS string) {
	cmd := killCMDgen(hOS)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errors.TMPerrorHandle(err)
}
