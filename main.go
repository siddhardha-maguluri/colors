package main

import "fmt"

type log struct{}

var TEXT_COLORS = map[string]string{
	"BLACK":  "\033[0m",
	"RED":    "\033[31m",
	"GREEN":  "\033[32m",
	"YELLOW": "\033[33m",
	"BLUE":   "\033[34m",
}

func (l log) Display(message string, displayColor string) {
	fmt.Printf("%s%s%s", TEXT_COLORS[displayColor], message, TEXT_COLORS["BLACK"])
}
