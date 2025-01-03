package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const AccountIDPrefix = "Account"

// AccountID is a value object for group chat id.
type AccountID struct {
	value string
}

// NewAccountID is a constructor for AccountID with generating id.
func NewAccountID() AccountID {
	id := uuid.Must(uuid.NewV7())
	return AccountID{value: id.String()}
}

// NewAccountIDFromString is a constructor for AccountID
// It creates AccountID from string
func NewAccountIDFromString(value string) (*AccountID, error) {
	if value == "" {
		return nil, errors.New("AccountID is empty")
	}
	if len(value) > len(AccountIDPrefix) && value[0:len(AccountIDPrefix)] == AccountIDPrefix {
		value = value[len(AccountIDPrefix)+1:]
	}
	return &AccountID{value: value}, nil
}

// ToJSON converts to JSON.
//
// However, this method is out of layer.
func (g *AccountID) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": g.value,
	}
}

func (g *AccountID) GetValue() string {
	return g.value
}

func (g *AccountID) GetTypeName() string {
	return AccountIDPrefix
}

func (g *AccountID) AsString() string {
	return fmt.Sprintf("%s-%s", g.GetTypeName(), g.GetValue())
}

func (g *AccountID) String() string {
	return g.AsString()
}

// Equals compares other AccountID.
func (g *AccountID) Equals(other *AccountID) bool {
	return g.value == other.value
}

// ConvertAccountIDFromJSON converts from JSON.
func ConvertAccountIDFromJSON(value map[string]interface{}) (*AccountID, error) {
	return NewAccountIDFromString(value["value"].(string))
}
