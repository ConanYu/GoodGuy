package util

import "runtime"

func GetFileName() string {
	_, fileName, _, _ := runtime.Caller(1)
	return fileName
}
