package reporting

import (
	"fmt"
	"os"
	"strings"
)

const newline = "\n"

var (
	success         = "✔"
	failure         = "✘"
	error_          = "🔥"
	skip            = "⚠"
	dotSuccess      = "."
	dotFailure      = "x"
	dotError        = "E"
	dotSkip         = "S"
	errorTemplate   = "* %s \nLine %d: - %v \n%s\n"
	failureTemplate = "* %s \nLine %d:\n%s\n"
)

var (
	greenColor  = "\033[32m"
	yellowColor = "\033[33m"
	redColor    = "\033[31m"
	resetColor  = "\033[0m"
)

func init() {
	if !xterm() {
		monochrome()
	}
}

// QuiteMode disables all console output symbols. This is only meant to be used
// for tests that are internal to goconvey where the output is distracting or
// otherwise not needed in the test output.
func QuietMode() {
	success, failure, error_, skip, dotSuccess, dotFailure, dotError, dotSkip = "", "", "", "", "", "", "", ""
}

func monochrome() {
	greenColor, yellowColor, redColor, resetColor = "", "", "", ""
}

func xterm() bool {
	return strings.Contains(fmt.Sprintf("%v", os.Environ()), " TERM=xterm")
}
