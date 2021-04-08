package main

import (
	"strings"

	"github.com/bitflipp/crude"
	"github.com/stoewer/go-strcase"
)

var (
	converters = map[string]crude.Converter{
		"kebab":      strcase.KebabCase,
		"lowerCamel": strcase.LowerCamelCase,
		"single": func(value string) string {
			if len(value) == 0 {
				return ""
			}
			return strings.ToLower(string(value[0]))
		},
		"snake":      strcase.SnakeCase,
		"upperCamel": strcase.UpperCamelCase,
		"upperKebab": strcase.UpperKebabCase,
		"upperSnake": strcase.UpperSnakeCase,
	}
)
