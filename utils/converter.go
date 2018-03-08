package utils

import (
	"fmt"
)

func ConvInterface2String(i interface{}) string {
	switch i.(type) {
	case string:
		return fmt.Sprintf("%v", i)
	default:
		return ""
	}
}
