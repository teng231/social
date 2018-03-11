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

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
