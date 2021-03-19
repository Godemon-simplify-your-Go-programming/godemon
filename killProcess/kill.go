package killProcess

import (
	"godemon/errors"
)

func KillProcess(hOS string) {
	cmd := killCMDgen(hOS)
	err := cmd.Run()
	errors.TMPerrorHandle(err)
}
