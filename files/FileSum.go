package files

import (
	"crypto/md5"
	"os"
	"io"
	"fmt"
	"errors"
)

func Md5Sum(path string) (md5sumStr string, myError error) {
	newHash := md5.New()
	file, myError := os.Open(path)
	if myError != nil {
		return
	}
	if isDir(file) {
		myError = errors.New("Given path \"" + path + "\" is a directory, not a file.")
		return
	}
	io.Copy(newHash, file)
	md5sumStr = fmt.Sprintf("%x",newHash.Sum(nil))
	return 
}

func isDir(file *os.File) bool {
	info, _ := file.Stat()
	return info.IsDir()
}