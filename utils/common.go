package utils

import (
	"encoding/json"
	"errors"
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

func Str2T(str string, T interface{}) error {
	if str == "" {
		e := errors.New("str empty")
		ErrLog(e)
		return e
	}
	err := json.Unmarshal([]byte(str), &T)
	if err != nil {
		ErrLog(err)
		return err
	}
	return nil
}

func ErrStr(err error) string {
	return fmt.Sprintf("%v", err)
}
