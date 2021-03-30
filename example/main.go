package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/bitflipp/crude"
)

func main() {
	es := make(map[string]*crude.Entity)
	f, err := os.Open("entities.json")
	if err != nil {
		log.Fatalf("failed to open entity definition file: %s", err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&es); err != nil {
		log.Fatalf("failed to decode entity definition file: %s", err)
	}
	t, err := template.New("crude.gohtml").Funcs(crude.FuncMap).ParseFiles("crude.gohtml")
	if err != nil {
		log.Fatalf("failed to parse template text: %s", err)
	}
	g := crude.Generator{Entities: es, Template: t}
	if err := g.Run(os.Stdout); err != nil {
		log.Fatalf("failed to run generator: %s", err)
	}
}
