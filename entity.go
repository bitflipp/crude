package crude

import (
	"errors"
	"fmt"
	"reflect"
)

type Entity struct {
	TableName    string
	ReceiverName string
	Value        interface{}
	Custom       interface{}

	reflectValue reflect.Value
}

func (e *Entity) validate() error {
	if e.TableName == "" {
		return errors.New("TableName is empty")
	}
	if e.ReceiverName == "" {
		return errors.New("ReceiverName is empty")
	}
	if e.Value == nil {
		return errors.New("Value is nil")
	}
	value := reflect.ValueOf(e.Value)
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("Value is not a struct")
	}
	e.reflectValue = value
	return nil
}
