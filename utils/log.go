package utils

import (
	"fmt"
	"log"
	"runtime"
)

const (
	DefaultLimit = 20
)

func Log(log interface{}) {
	fmt.Printf("%v", log)
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
