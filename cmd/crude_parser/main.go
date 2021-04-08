package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/bitflipp/crude"
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
	tableConverterNames := ""
	flag.StringVar(&flags.InputFilePath, "i", "", "Input file path. If empty, stdin is used")
	flag.StringVar(&flags.OutputFilePath, "o", "", "Output file path. If empty, stdout is used")
	flag.StringVar(&flags.ReceiverConverter, "rc", "single", "Receiver converter, choice of "+tableConverterNames)
	flag.StringVar(&flags.TableConverter, "tc", "snake", "Table converter, choice of "+tableConverterNames)
	flag.StringVar(&flags.ColumnConverter, "fc", "snake", "Column converter, choice of "+tableConverterNames)
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
	output := os.Stdin
	if flags.OutputFilePath != "" {
		var err error
		output, err = os.Create(flags.OutputFilePath)
		if err != nil {
			log.Fatalf("failed to open output file: %s", err)
		}
		defer output.Close()
	}

	// Parser
	receiverConverter, ok := converters[flags.ReceiverConverter]
	if !ok {
		log.Fatalf("unknown receiver converter: %s", flags.ReceiverConverter)
	}
	tableConverter, ok := converters[flags.TableConverter]
	if !ok {
		log.Fatalf("unknown table converter: %s", flags.TableConverter)
	}
	columnConverter, ok := converters[flags.ColumnConverter]
	if !ok {
		log.Fatalf("unknown field converter: %s", flags.ColumnConverter)
	}
	parser := crude.Parser{
		Input:             input,
		FileName:          fileName,
		ReceiverConverter: receiverConverter,
		TableConverter:    tableConverter,
		ColumnConverter:   columnConverter,
	}
	if err := parser.Run(output); err != nil {
		log.Fatalf("failed to run parser: %s", err)
	}
}