package zlog

import "strings"

type Output int8

const (
	Stderr Output = iota
	Stdout
	File
)

func (o Output) String() string {
	switch o {
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

func ParseOutput(s string) Output {
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
