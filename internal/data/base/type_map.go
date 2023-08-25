package base

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TypeMap map[string]interface{}

func NewTypeMap(value map[string]interface{}) TypeMap {
	return value
}

// Value Marshal
func (a TypeMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *TypeMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func (a TypeMap) ToMap() map[string]interface{} {
	var result = make(map[string]interface{})
	for k, v := range a {
		result[k] = v
	}
	return result
}
