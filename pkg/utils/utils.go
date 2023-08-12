package utils

import "path/filepath"

func GetFileName(fileName string) string {
	file := filepath.Base(fileName)

	return file[:len(file)-len(filepath.Ext(file))]
}