package controllers

func WatchFiles(fileordirPath string, hOS string, cnf string) error {
	return watch(fileordirPath, hOS, cnf)
}
