package crude

import (
	"fmt"
)

type definition struct {
	TableName    string
	ReceiverName string
	EntityName   string
	FieldNames   []string
	ColumnNames  []string
	Custom       interface{}

	fieldNames  map[string]string
	columnNames map[string]string
}

func (d *definition) ToFieldNames(columnNames []string) []string {
	fieldNames := make([]string, len(columnNames))
	for i, columnName := range columnNames {
		fieldNames[i] = d.fieldNames[columnName]
	}
	return fieldNames
}

func (d *definition) ToColumnNames(fieldNames []string) []string {
	columnNames := make([]string, len(fieldNames))
	for i, fieldName := range fieldNames {
		columnNames[i] = d.columnNames[fieldName]
	}
	return columnNames
}

func (d *definition) derive(entity Entity) error {
	entityName := entity.reflectValue.Type().Name()
	d.fieldNames = make(map[string]string)
	d.columnNames = make(map[string]string)
	for i := 0; i < entity.reflectValue.NumField(); i++ {
		fieldType := entity.reflectValue.Type().Field(i)
		fieldName := fieldType.Name
		columnName := fieldType.Tag.Get("crude")
		if columnName == "" {
			continue
		}
		for _, existingColumnName := range d.columnNames {
			if columnName == existingColumnName {
				return fmt.Errorf("duplicate column name in entity '%s': '%s'", entityName, columnName)
			}
		}
		d.fieldNames[columnName] = fieldName
		d.columnNames[fieldName] = columnName
	}
	d.EntityName = entityName
	d.ReceiverName = entity.ReceiverName
	d.TableName = entity.TableName
	d.FieldNames = make([]string, len(d.columnNames))
	d.ColumnNames = make([]string, len(d.columnNames))
	i := 0
	for fieldName, columnName := range d.columnNames {
		d.FieldNames[i] = fieldName
		d.ColumnNames[i] = columnName
		i++
	}
	d.Custom = entity.Custom
	return nil
}
