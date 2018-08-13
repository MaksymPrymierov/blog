package routes

import (
	"html/template"
)

/* Function for html template */
func unescape(x string) interface{} {
	return template.HTML(x)
}

func checkGroup(group string) bool {
	if group != "admin" {
		return false
	}
	return true
}
