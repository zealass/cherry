package cherryFile

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func JudgePath(filePath string) (string, bool) {
	tmpPath := path.Join(GetMainFuncDir(), filePath)
	ok := IsDir(tmpPath)
	if ok {
		return tmpPath, true
	}

	tmpPath = path.Join(GetWorkPath(), filePath)
	ok = IsDir(tmpPath)
	if ok {
		return tmpPath, true
	}

	ok = IsDir(filePath)
	if ok {
		return filePath, true
	}

	return "", false
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true
	}
	return false
}

var (
	mainFuncDir = ""
)

func GetMainFuncDir() string {
	if mainFuncDir == "" {
		var buf [2 << 16]byte
		stack := string(buf[:runtime.Stack(buf[:], true)])

		lines := strings.Split(strings.TrimSpace(stack), "\n")
		lastLine := strings.TrimSpace(lines[len(lines)-1])

		lastIndex := strings.LastIndex(lastLine, "/")
		if lastIndex < 1 {
			return ""
		}

		mainFuncDir = lastLine[:lastIndex]
	}
	return mainFuncDir
}

func GetWorkPath() string {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return p
}

func JoinPath(elem ...string) (string, error) {
	filePath := path.Join(elem...)

	err := CheckPath(filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func CheckPath(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	return err
}

func GetFileName(filePath string, removeExt bool) string {
	fileName := path.Base(filePath)
	if removeExt == false {
		return fileName
	}

	var suffix string
	suffix = path.Ext(fileName)

	return strings.TrimSuffix(fileName, suffix)
}
