package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const PymonIDPrefix = "Pymon"

// PymonID is a value object for group chat id.
type PymonID struct {
	value string
}

// NewPymonID is a constructor for PymonID with generating id.
func NewPymonID() PymonID {
	id := uuid.Must(uuid.NewV7())
	return PymonID{value: id.String()}
}

// NewPymonIDFromString is a constructor for PymonID
// It creates PymonID from string
func NewPymonIDFromString(value string) (*PymonID, error) {
	if value == "" {
		return nil, errors.New("PymonID is empty")
	}
	if len(value) > len(PymonIDPrefix) && value[0:len(PymonIDPrefix)] == PymonIDPrefix {
		value = value[len(PymonIDPrefix)+1:]
	}
	return &PymonID{value: value}, nil
}

// ToJSON converts to JSON.
//
// However, this method is out of layer.
func (g *PymonID) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": g.value,
	}
}

func (g *PymonID) GetValue() string {
	return g.value
}

func (g *PymonID) GetTypeName() string {
	return PymonIDPrefix
}

func (g *PymonID) AsString() string {
	return fmt.Sprintf("%s-%s", g.GetTypeName(), g.GetValue())
}

func (g *PymonID) String() string {
	return g.AsString()
}

// Equals compares other PymonID.
func (g *PymonID) Equals(other *PymonID) bool {
	return g.value == other.value
}

// ConvertPymonIDFromJSON converts from JSON.
func ConvertPymonIDFromJSON(value map[string]interface{}) (*PymonID, error) {
	return NewPymonIDFromString(value["value"].(string))
}
