package utils

import (
	"os"
	"path"
	"strings"
)

// FileGetCurrentDirectory 获取当前目录
func FileGetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return dir
}

// FileGetParentDirectory 获取上级目录
func FileGetParentDirectory(directory string) string {
	return StringSub(directory, 0, strings.LastIndex(directory, "/"))
}

// FileGetName 获取文件名称
func FileGetName(filePath string) string {
	return path.Base(filePath)
}

// FileGetSize 获取取文件大小
func FileGetSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	fileSize := fileInfo.Size()
	return fileSize, nil
}

// FileExists 文件是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}