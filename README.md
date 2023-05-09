# Colorful Logging in Go

This is a simple module that allows you to print text in the color of your choice in Go.

## Installation

You can install this package by running the following command:

`go get github.com/siddhardha-maguluri/colors`

## Usage

To use this package, you need to import it in your Go code:

`import "github.com/siddhardha-maguluri/colors"`

You can then use the logger.Warn() function to print warning messages in red:

`logger.Warn("This is a warning message!")`

You can also use the logger.Info() function to print info messages in blue:

`logger.Info("This is an info message!")`

## Customization

You can customize the colors used by this package by modifying the TEXT_COLORS map in colors.go. Here are the default colors:

<code>
TEXT_COLORS["BLACK"] = "\033[0m"

TEXT_COLORS["RED"] = "\033[31m"

TEXT_COLORS["GREEN"] = "\033[32m"

TEXT_COLORS["YELLOW"] = "\033[33m"

TEXT_COLORS["BLUE"] = "\033[34m"
</code>

You can add or remove colors as needed.