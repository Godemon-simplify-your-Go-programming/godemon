package controllers

func WatchFiles(fileordirPath string, hOS string) error {
	return watch(fileordirPath, hOS)
}
