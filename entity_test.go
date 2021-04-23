package crude

import (
	"reflect"
	"testing"
)

func newTestEntity() Entity {
	return Entity{
		Name:         "entity",
		Table:        "entity",
		Receiver:     "e",
		InsertFields: make([]string, 0),
		FieldColumns: map[string]string{
			"Field1": "column1",
			"field2": "column2",
		},
	}
}

func TestInvalidEntity(t *testing.T) {
	entity := Entity{}
	if err := entity.validate(); err == nil {
		t.Error("empty Name was accepted")
	}
	entity.Name = "entity"
	if err := entity.validate(); err == nil {
		t.Error("empty Table was accepted")
	}
	entity.Table = "entity"
	if err := entity.validate(); err == nil {
		t.Error("empty Receiver was accepted")
	}
	entity.Receiver = "e"
	if err := entity.validate(); err == nil {
		t.Error("nil InsertFields was accepted")
	}
	entity.InsertFields = make([]string, 0)
	if err := entity.validate(); err == nil {
		t.Error("nil FieldColumns was accepted")
	}
	entity.FieldColumns = make(map[string]string)
	if err := entity.validate(); err == nil {
		t.Error("empty FieldColumns was accepted")
	}
	entity.FieldColumns["Field"] = "column"
	if err := entity.validate(); err != nil {
		t.Errorf("failed to validate valid Entity: %s", err)
	}
}

func TestEntityToFields(t *testing.T) {
	entity := newTestEntity()
	fields := entity.ToFields([]string{"column1", "column2"})
	expectedFields := []string{"Field1", "field2"}
	if !reflect.DeepEqual(fields, expectedFields) {
		t.Errorf("got: %s, expected: %s", fields, expectedFields)
	}
}

func TestEntityToColumns(t *testing.T) {
	entity := newTestEntity()
	columns := entity.ToColumns([]string{"Field1", "field2"})
	expectedColumns := []string{"column1", "column2"}
	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Errorf("got: %s, expected: %s", columns, expectedColumns)
	}
}
