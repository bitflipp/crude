package main

import (
	"log"
	"os"
	"text/template"

	"github.com/bitflipp/crude"
)

type author struct {
	ID   int64  `crude:"id"`
	Name string `crude:"name"`
}

type book struct {
	ID     int64  `crude:"id"`
	Name   string `crude:"name"`
	ISBN   string `crude:"isbn"`
	Author *author
}

func main() {
	t, err := template.New("crude.gohtml").Funcs(crude.FuncMap).ParseFiles("crude.gohtml")
	if err != nil {
		log.Fatalf("failed to parse template text: %s", err)
	}
	g := crude.Generator{
		Entities: []crude.Entity{
			{TableName: "author", ReceiverName: "a", Value: author{}, Custom: []string{"Name"}},
			{TableName: "book", ReceiverName: "b", Value: book{}, Custom: []string{"Name", "ISBN"}},
		},
		Template: t,
	}
	if err := g.Run(os.Stdout); err != nil {
		log.Fatalf("failed to run generator: %s", err)
	}
}
