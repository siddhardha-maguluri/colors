package logger

import "fmt"

type log struct{}

func (l log) Error(message string) {
	fmt.Printf("%s%s%s", TEXT_COLORS["RED"], message, TEXT_COLORS["BLACK"])
}

func (l log) Warn(message string) {
	fmt.Printf("%s%s%s", TEXT_COLORS["YELLOW"], message, TEXT_COLORS["BLACK"])
}

func (l log) Info(message string) {
	fmt.Printf("%s%s%s", TEXT_COLORS["BLUE"], message, TEXT_COLORS["BLACK"])
}
