package utils

import (
	"log"
	"testing"
)

func TestStr2T(t *testing.T) {
	str := "[\"absc\",\"cde\"]"
	var a []string
	err := Str2T(str, &a)
	log.Printf("%v", err)
	log.Printf("%v", a)
	if len(a) == 0 || err != nil {
		log.Fatalf("%s: %v", "err", err)
	}

	type T struct {
		User string `json:"user"`
	}
	str2 := "[{\"user\":\"abc\"},{\"user\":\"cde\"}]"
	var a2 []T
	err2 := Str2T(str2, &a2)
	log.Printf("%v", a2[0].User)
	if len(a2) == 0 || err2 != nil {
		log.Fatalf("%s: %v", "err", err2)
	}
}
