package crude

import (
	"errors"
	"fmt"
	"io"
	"text/template"
)

type Generator struct {
	Entities map[string]*Entity
	Template *template.Template
}

func (g *Generator) validate() error {
	if g.Entities == nil || len(g.Entities) == 0 {
		return errors.New("Entities is nil or empty")
	}
	if g.Template == nil {
		return errors.New("Template is nil")
	}
	for name, entity := range g.Entities {
		entity.Name = name
		if err := entity.validate(); err != nil {
			return fmt.Errorf("failed to validate entity: %w", err)
		}
	}
	return nil
}

func (g *Generator) Run(writer io.Writer) error {
	if err := g.validate(); err != nil {
		return fmt.Errorf("failed to validate generator: %w", err)
	}
	return g.Template.Execute(writer, g.Entities)
}
