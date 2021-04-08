package main

import (
	"strings"

	"github.com/bitflipp/crude"
	"github.com/stoewer/go-strcase"
)

var (
	converters = map[string]crude.Converter{
		"single": func(value string) string {
			if len(value) == 0 {
				return ""
			}
			return strings.ToLower(string(value[0]))
		},
		"kebab":      strcase.KebabCase,
		"lowerCamel": strcase.LowerCamelCase,
		"snake":      strcase.SnakeCase,
		"upperCamel": strcase.UpperCamelCase,
		"upperKebab": strcase.UpperKebabCase,
		"upperSnake": strcase.UpperSnakeCase,
	}
)
