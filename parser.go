package crude

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type visitor struct {
	parser   *Parser
	entities map[string]Entity
	typeSpec *ast.TypeSpec
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch nodeImpl := node.(type) {
	case *ast.TypeSpec:
		v.typeSpec = nodeImpl
	case *ast.StructType:
		if v.typeSpec == nil {
			return v
		}
		typeName := v.typeSpec.Name.Name
		entity := Entity{
			Name:         typeName,
			Table:        v.parser.TableConverter(typeName),
			Receiver:     v.parser.ReceiverConverter(typeName),
			InsertFields: make([]string, 0),
			FieldColumns: make(map[string]string),
		}
		for _, field := range nodeImpl.Fields.List {
			fieldName := field.Names[0].Name
			columnName := v.parser.ColumnConverter(fieldName)
			entity.FieldColumns[fieldName] = columnName
			entity.InsertFields = append(entity.InsertFields, fieldName)
		}
		v.entities[entity.Name] = entity
		v.typeSpec = nil
	}
	return v
}

type Converter func(string) string

type Parser struct {
	Input             io.Reader
	FileName          string
	ReceiverConverter Converter
	TableConverter    Converter
	ColumnConverter   Converter
}

func (p *Parser) Run(output io.Writer) error {
	src, err := ioutil.ReadAll(p.Input)
	if err != nil {
		return err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, p.FileName, src, 0)
	if err != nil {
		return err
	}
	visitor := &visitor{
		parser:   p,
		entities: make(map[string]Entity),
	}
	ast.Walk(visitor, file)
	return toml.NewEncoder(output).Encode(visitor.entities)
}
