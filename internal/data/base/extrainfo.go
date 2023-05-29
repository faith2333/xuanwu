package base

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ExtraInfo map[string]interface{}

// Value Marshal
func (a ExtraInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *ExtraInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
