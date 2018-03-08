package utils

import (
	"fmt"
)

const (
	DefaultLimit = 20
)

func Log(log interface{}) {
	fmt.Printf("%v", log)
	fmt.Println()
}
