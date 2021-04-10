package main

import (
	"flag"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/bitflipp/crude"
	"github.com/pelletier/go-toml"
)

var (
	flags struct {
		InputFilePath     string
		OutputFilePath    string
		ReceiverConverter string
		TableConverter    string
		ColumnConverter   string
	}
)

func main() {
	// Flags
	converterNames := make(sort.StringSlice, 0)
	for converterName := range converters {
		converterNames = append(converterNames, converterName)
	}
	converterNames.Sort()
	converterNamesJoined := strings.Join(converterNames, ", ")
	flag.StringVar(&flags.ColumnConverter, "cc", "snake", "Column converter, choice of "+converterNamesJoined)
	flag.StringVar(&flags.InputFilePath, "i", "", "Input file path. If empty, stdin is used")
	flag.StringVar(&flags.OutputFilePath, "o", "", "Output file path. If empty, stdout is used")
	flag.StringVar(&flags.ReceiverConverter, "rc", "single", "Receiver converter, choice of "+converterNamesJoined)
	flag.StringVar(&flags.TableConverter, "tc", "snake", "Table converter, choice of "+converterNamesJoined)
	flag.Parse()

	// Input
	input := os.Stdin
	fileName := "src.go"
	if flags.InputFilePath != "" {
		var err error
		input, err = os.Open(flags.InputFilePath)
		if err != nil {
			log.Fatalf("failed to open input file: %s", err)
		}
		defer input.Close()
		fileName = path.Base(flags.InputFilePath)
	}

	// Output
	output := os.Stdout
	if flags.OutputFilePath != "" {
		var err error
		output, err = os.Create(flags.OutputFilePath)
		if err != nil {
			log.Fatalf("failed to open output file: %s", err)
		}
		defer output.Close()
	}

	// Parser
	receiverConverter, found := converters[flags.ReceiverConverter]
	if !found {
		log.Fatalf("unknown receiver converter: %s", flags.ReceiverConverter)
	}
	tableConverter, found := converters[flags.TableConverter]
	if !found {
		log.Fatalf("unknown table converter: %s", flags.TableConverter)
	}
	columnConverter, found := converters[flags.ColumnConverter]
	if !found {
		log.Fatalf("unknown column converter: %s", flags.ColumnConverter)
	}
	parser := crude.Parser{
		Input:             input,
		FileName:          fileName,
		ReceiverConverter: receiverConverter,
		TableConverter:    tableConverter,
		ColumnConverter:   columnConverter,
	}
	entities, err := parser.Run()
	if err != nil {
		log.Fatalf("failed to run parser: %s", err)
	}
	if err := toml.NewEncoder(output).Encode(entities); err != nil {
		log.Fatalf("failed to encode entities: %s", err)
	}
}
