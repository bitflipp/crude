package crude

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func identityConverter(value string) string {
	return value
}

func newTestParser(input io.Reader) *Parser {
	return &Parser{
		Input:             input,
		FileName:          "test.go",
		ReceiverConverter: identityConverter,
		TableConverter:    identityConverter,
		ColumnConverter:   identityConverter,
	}
}

func parseEntity(source string, entityName string) (*Parser, *Entity, error) {
	parser := newTestParser(bytes.NewBufferString(source))
	entities, err := parser.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run parser: %w", err)
	}
	if len(entities) != 1 {
		return nil, nil, errors.New("number of entities is not 1")
	}
	entity, found := entities[entityName]
	if !found {
		return nil, nil, fmt.Errorf("failed to find entity '%s'", entityName)
	}
	if err := entity.validate(); err != nil {
		return nil, nil, fmt.Errorf("failed to validate entity: %s", err)
	}
	return parser, &entity, nil
}

func TestStructWithMixedFields(t *testing.T) {
	source := `
package test

type entity struct {
	A int
	B string
	c bool
}
`
	entityName := "entity"
	parser, entity, err := parseEntity(source, entityName)
	if err != nil {
		t.Errorf("failed to parse entity: %s", err)
	}
	tableExpected := parser.TableConverter(entityName)
	if entity.Table != tableExpected {
		t.Errorf("Table: got %s, expected %s", entity.Table, tableExpected)
	}
	receiverExpected := parser.ReceiverConverter(entityName)
	if entity.Receiver != receiverExpected {
		t.Errorf("Receiver: got %s, expected %s", entity.Receiver, receiverExpected)
	}
	columnFieldsExpected := map[string]string{
		"A": parser.ColumnConverter("A"),
		"B": parser.ColumnConverter("B"),
		"c": parser.ColumnConverter("c"),
	}
	if !reflect.DeepEqual(entity.FieldColumns, columnFieldsExpected) {
		t.Errorf("ColumnFields: got %#v, expected %#v", entity.ColumnFields, columnFieldsExpected)
	}
}

func TestUnrelatedAnonymousStruct(t *testing.T) {
	source := `
package test

type entity struct {
	A int
	B string
	c bool
}

func anonymous() {
	a := struct{
		D float64
	}{}
}
`
	entityName := "entity"
	_, entity, err := parseEntity(source, entityName)
	if err != nil {
		t.Errorf("failed to parse entity: %s", err)
	}
	if _, found := entity.FieldColumns["D"]; found {
		t.Error("entity contains field from unrelated anonymous struct")
	}
}
