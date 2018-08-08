package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	fmt.Printf(string(b))
	return fmt.Sprintf("%x", b)
}

func GenerateNameId(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "/", "", -1)
	s = strings.Replace(s, "\\", "", -1)
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	s = strings.Replace(s, ":", "", -1)
	s = strings.Replace(s, ";", "", -1)
	s = strings.Replace(s, "|", "", -1)
	s = strings.Replace(s, "=", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "+", "", -1)
	s = strings.Replace(s, "*", "", -1)
	s = strings.Replace(s, "?", "", -1)
	s = strings.Replace(s, "<", "", -1)
	s = strings.Replace(s, ">", "", -1)
	s = strings.Replace(s, " ", "-", -1)
	return s
}
