package crude

import (
	"errors"
	"fmt"
	"io"
	"text/template"
)

type Generator struct {
	Entities []Entity
	Template *template.Template
}

func (g *Generator) validate() error {
	if g.Entities == nil || len(g.Entities) == 0 {
		return errors.New("Entities is nil or empty")
	}
	if g.Template == nil {
		return errors.New("Template is nil")
	}
	return nil
}

func (g *Generator) Run(writer io.Writer) error {
	if err := g.validate(); err != nil {
		return fmt.Errorf("failed to validate generator: %w", err)
	}
	definitions := make(map[string]*definition)
	for _, entity := range g.Entities {
		definition := &definition{}
		if err := entity.validate(); err != nil {
			return fmt.Errorf("failed to validate entity: %w", err)
		}
		if err := definition.derive(entity); err != nil {
			return fmt.Errorf("failed to derive table definition: %w", err)
		}
		if _, found := definitions[definition.TableName]; found {
			return fmt.Errorf("duplicate table name: '%s'", definition.TableName)
		}
		definitions[definition.TableName] = definition
	}
	return g.Template.Execute(writer, definitions)
}
