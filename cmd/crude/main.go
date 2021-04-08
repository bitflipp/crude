package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/bitflipp/crude"
	"github.com/pelletier/go-toml"
)

var (
	flags struct {
		EntityDefinitionFilePath string
		TemplateFilePath         string
		OutputFilePath           string
	}
)

func main() {
	// Flags
	flag.StringVar(&flags.EntityDefinitionFilePath, "e", "./crude.toml", "Entity definition file path")
	flag.StringVar(&flags.TemplateFilePath, "t", "./crude.gohtml", "Template file path")
	flag.StringVar(&flags.OutputFilePath, "o", "", "Output file path. If empty, stdout is used")
	flag.Parse()

	// Entities
	entities := make(map[string]*crude.Entity)
	entityDefinitionFile, err := os.Open(flags.EntityDefinitionFilePath)
	if err != nil {
		log.Fatalf("failed to open entity definition file: %s", err)
	}
	defer entityDefinitionFile.Close()
	if err := toml.NewDecoder(entityDefinitionFile).Decode(&entities); err != nil {
		log.Fatalf("failed to decode entity definition file: %s", err)
	}

	// Templates
	crudeTemplates, err := template.New("crude.gohtml").Funcs(crude.FuncMap).ParseFiles(flags.TemplateFilePath)
	if err != nil {
		log.Fatalf("failed to parse template file: %s", err)
	}

	// Generator
	generator := crude.Generator{Entities: entities, Template: crudeTemplates}
	outputFile := os.Stdout
	if flags.OutputFilePath != "" {
		var err error
		outputFile, err = os.Create(flags.OutputFilePath)
		if err != nil {
			log.Fatalf("failed to create output file: %s", err)
		}
		defer outputFile.Close()
	}
	if err := generator.Run(outputFile); err != nil {
		log.Fatalf("failed to run generator: %s", err)
	}
}
