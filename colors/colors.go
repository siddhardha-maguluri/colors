package colors

import "fmt"

type Log struct{}

// Use to display the error messages in red color
func (l *Log) Error(message string) {
	fmt.Printf("%s%s%s", TEXT_COLORS["RED"], message, TEXT_COLORS["BLACK"])
	fmt.Println()
}

// Use to display the warning messages in yellow color
func (l *Log) Warn(message string) {
	fmt.Printf("%s%s%s", TEXT_COLORS["YELLOW"], message, TEXT_COLORS["BLACK"])
	fmt.Println()
}

// Use to display the info messages in blue color
func (l *Log) Info(message string) {
	fmt.Printf("%s%s%s", TEXT_COLORS["BLUE"], message, TEXT_COLORS["BLACK"])
	fmt.Println()
}
