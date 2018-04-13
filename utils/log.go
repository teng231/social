package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

const (
	DefaultLimit = 20
)

func Log(logs ...interface{}) {
	for _, log := range logs {
		fmt.Printf("%v", log)
	}
	fmt.Println()
}

func ErrLog(err error) (b bool) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, err)
		b = true
	}
	return
}

func ReadFileRoot(path string) ([]byte, error) {
	if &path == nil || path == "" {
		return nil, errors.New("no path")
	}
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		ErrLog(err)
	}
	return data, err
}

func Jsonify(j interface{}) string {
	bytes, err := json.Marshal(j)
	ErrLog(err)
	return string(bytes[:])
}

func LogJson(j interface{}) {
	fmt.Println(Jsonify(j))
}
