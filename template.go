package crude

import (
	"strings"
	"text/template"
)

func join(separator string, values []string) string {
	return strings.Join(values, separator)
}

func repeat(times int, value string) []string {
	repeated := make([]string, times)
	for i := 0; i < times; i++ {
		repeated[i] = value
	}
	return repeated
}

func wrap(left, right string, values []string) []string {
	newValues := make([]string, len(values))
	for i, value := range values {
		newValues[i] = left + value + right
	}
	return newValues
}

func zip(separator string, leftValues []string, rightValues []string) []string {
	if len(leftValues) != len(rightValues) {
		return nil
	}
	zipped := make([]string, len(leftValues))
	for i := 0; i < len(leftValues); i++ {
		zipped[i] = leftValues[i] + separator + rightValues[i]
	}
	return zipped
}

var (
	FuncMap = template.FuncMap{
		"join":   join,
		"repeat": repeat,
		"wrap":   wrap,
		"zip":    zip,
	}
)
