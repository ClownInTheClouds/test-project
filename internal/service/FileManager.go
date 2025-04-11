package service

import (
	"log"
	"os"
	"path/filepath"
)

func GetFile(path ...string) (*os.File, error) {
	var exePath = getExePath()
	var exeDir = filepath.Dir(exePath)
	var targetPath = filepath.Join(append([]string{exeDir}, path...)...)

	file, err := os.OpenFile(targetPath, os.O_RDWR, 0644)
	if err == nil {
		return file, nil
	}

	if os.IsNotExist(err) {
		dir := filepath.Dir(targetPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}

		file, err = os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return nil, err
}

func getExePath() string {
	var exePath, err = os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return exePath
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
