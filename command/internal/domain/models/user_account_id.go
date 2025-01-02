package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const UserAccountIdPrefix = "UserAccount"

// UserAccountId is a value object that represents a user account id.
type UserAccountId struct {
	value string
}

// NewUserAccountId is the constructor for UserAccountId with generating id.
func NewUserAccountId() UserAccountId {
	id := uuid.Must(uuid.NewV7())
	return UserAccountId{value: id.String()}
}

// NewUserAccountIdFromString is the constructor for UserAccountId.
func NewUserAccountIdFromString(value string) (*UserAccountId, error) {
	if value == "" {
		return nil, errors.New("UserAccountId is empty")
	}
	if len(value) > len(UserAccountIdPrefix) && value[0:len(UserAccountIdPrefix)] == UserAccountIdPrefix {
		value = value[len(UserAccountIdPrefix)+1:]
	}
	return &UserAccountId{value: value}, nil
}

// ConvertUserAccountIdFromJSON is a constructor for UserAccountId.
func ConvertUserAccountIdFromJSON(value map[string]interface{}) (*UserAccountId, error) {
	return NewUserAccountIdFromString(value["value"].(string))
}

// ToJSON converts to JSON.
//
// However, this method is out of layer.
func (u *UserAccountId) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": u.value,
	}
}

func (u *UserAccountId) GetValue() string {
	return u.value
}

func (u *UserAccountId) GetTypeName() string {
	return "UserAccount"
}

func (u *UserAccountId) AsString() string {
	return fmt.Sprintf("%s-%s", u.GetTypeName(), u.GetValue())
}

func (u *UserAccountId) String() string {
	return u.AsString()
}

func (u *UserAccountId) Equals(other *UserAccountId) bool {
	return u.value == other.value
}
