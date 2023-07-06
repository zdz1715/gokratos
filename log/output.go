package log

import "strings"

type Output int8

const (
	Stderr Output = iota
	Stdout        = iota
	File
)

func (l Output) String() string {
	switch l {
	case Stderr:
		return "stderr"
	case Stdout:
		return "stdout"
	case File:
		return "file"
	default:
		return ""
	}
}

func ParseLevel(s string) Output {
	switch strings.ToLower(s) {
	case "stderr":
		return Stderr
	case "stdout":
		return Stdout
	case "file":
		return File
	}
	return Stderr
}
