package crude

import (
	"errors"
)

type Entity struct {
	// Required
	Name         string
	Table        string
	Receiver     string
	InsertFields []string
	FieldColumns map[string]string

	// Optional
	Custom interface{}

	// Computed
	Fields       []string          `toml:"-"`
	ColumnFields map[string]string `toml:"-"`
	Columns      []string          `toml:"-"`
}

func (e *Entity) ToFields(columns []string) []string {
	fields := make([]string, len(columns))
	for i, column := range columns {
		fields[i] = e.ColumnFields[column]
	}
	return fields
}

func (e *Entity) ToColumns(fields []string) []string {
	columns := make([]string, len(fields))
	for i, field := range fields {
		columns[i] = e.FieldColumns[field]
	}
	return columns
}

func (e *Entity) validate() error {
	if e.Name == "" {
		return errors.New("Name is empty")
	}
	if e.Table == "" {
		return errors.New("Table is empty")
	}
	if e.Receiver == "" {
		return errors.New("Receiver is empty")
	}
	if e.InsertFields == nil {
		return errors.New("InsertFields is nil")
	}
	if e.FieldColumns == nil || len(e.FieldColumns) == 0 {
		return errors.New("FieldColumns is nil or empty")
	}
	e.Fields = make([]string, len(e.FieldColumns))
	e.ColumnFields = make(map[string]string)
	e.Columns = make([]string, len(e.FieldColumns))
	i := 0
	for fieldName, columnName := range e.FieldColumns {
		e.Fields[i] = fieldName
		e.ColumnFields[columnName] = fieldName
		e.Columns[i] = columnName
		i++
	}
	return nil
}
