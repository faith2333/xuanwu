package base

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TypeSlice []string

func NewTypeSlice(value []string) TypeSlice {
	return value
}

// Value Marshal
func (a TypeSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *TypeSlice) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func (a TypeSlice) ToSlice() []string {
	return a
}
