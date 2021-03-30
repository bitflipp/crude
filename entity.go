package crude

import (
	"errors"
)

type Entity struct {
	Name         string            `json:"-"`
	Table        string            `json:"table"`
	Receiver     string            `json:"receiver"`
	FieldColumns map[string]string `json:"fieldColumns"`
	Fields       []string          `json:"-"`
	ColumnFields map[string]string `json:"-"`
	Columns      []string          `json:"-"`
	Custom       []string          `json:"custom"`
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
		return errors.New("TableName is empty")
	}
	if e.Receiver == "" {
		return errors.New("ReceiverName is empty")
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
