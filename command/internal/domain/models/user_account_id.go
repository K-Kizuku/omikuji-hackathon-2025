package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const UserAccountIDPrefix = "UserAccount"

// UserAccountID is a value object that represents a user account id.
type UserAccountID struct {
	value string
}

// NewUserAccountID is the constructor for UserAccountID with generating id.
func NewUserAccountID() UserAccountID {
	id := uuid.Must(uuid.NewV7())
	return UserAccountID{value: id.String()}
}

// NewUserAccountIDFromString is the constructor for UserAccountID.
func NewUserAccountIDFromString(value string) (*UserAccountID, error) {
	if value == "" {
		return nil, errors.New("UserAccountID is empty")
	}
	if len(value) > len(UserAccountIDPrefix) && value[0:len(UserAccountIDPrefix)] == UserAccountIDPrefix {
		value = value[len(UserAccountIDPrefix)+1:]
	}
	return &UserAccountID{value: value}, nil
}

// ConvertUserAccountIDFromJSON is a constructor for UserAccountID.
func ConvertUserAccountIDFromJSON(value map[string]interface{}) (*UserAccountID, error) {
	return NewUserAccountIDFromString(value["value"].(string))
}

// ToJSON converts to JSON.
//
// However, this method is out of layer.
func (u *UserAccountID) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": u.value,
	}
}

func (u *UserAccountID) GetValue() string {
	return u.value
}

func (u *UserAccountID) GetTypeName() string {
	return "UserAccount"
}

func (u *UserAccountID) AsString() string {
	return fmt.Sprintf("%s-%s", u.GetTypeName(), u.GetValue())
}

func (u *UserAccountID) String() string {
	return u.AsString()
}

func (u *UserAccountID) Equals(other *UserAccountID) bool {
	return u.value == other.value
}
