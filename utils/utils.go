package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

/* Func return randome 16 bits numbers in string */
func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	fmt.Printf(string(b))
	return fmt.Sprintf("%x", b)
}

/* Func generate id of text */
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

// Function return true if string >= min or <= max
func CheckLen(s string, min, max int) bool {
	lenString := len(s)

	if lenString >= min && lenString <= max {
		return true
	} else {
		return false
	}
}
