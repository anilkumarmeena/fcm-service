package utils

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kataras/golog"
)

var httpDefaultClient = &http.Client{Timeout: 40 * time.Second}

func GetEnv(key, fallback string) (value string) {
	value = fallback
	if val, ok := os.LookupEnv(key); ok {
		value = val
	}
	return
}
func PathIs(name []byte, r *http.Request) bool {
	return bytes.Equal([]byte(r.URL.Path), name)
}

func SetKeyForQueryMap(query map[string][]string, key string, value string) {
	query[key] = []string{value}
}

func GetValueFromQueryMap(query map[string][]string, key string) string {
	if keyArray, ok := query[key]; ok {
		if len(keyArray) >= 1 {
			return keyArray[0]
		}
	}
	return ""
}

type responseBody struct {
	Msg  string `json:"msg,omitempty"`
	Code int    `json:"code,omitempty"`
}

func Mkdir(path string) bool {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash > 0 {
		path = path[:lastSlash]
	}
	var hasError = false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	} else if err != nil {
		golog.Error(err)
		hasError = true
	}
	return hasError
}

func ShellExec(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
func ArraySearch(list []string, key string) int {
	if list == nil {
		return -1
	}
	for i, item := range list {
		if item == key {
			return i
		}
	}
	return -1
}

func SplitDirAndFilename(path string) (string, string) {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		return "", path
	}
	return path[:lastSlash], path[lastSlash+1:]
}

func Now(ms bool) (ts int64) {
	ts = time.Now().UnixNano() / time.Millisecond.Nanoseconds()
	if !ms {
		ts = ts / 1000
	}
	return
}
