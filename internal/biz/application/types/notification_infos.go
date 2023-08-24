package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NotificationInfos []*NotificationInfo

type NotificationInfo struct {
	NotifyUsers []string `json:"notifyUsers"`
	// the id of the notify config in notification center.
	NotifyID string `json:"notifyId"`
}

// Value Marshal
func (infos NotificationInfos) Value() (driver.Value, error) {
	return json.Marshal(infos)
}

// Scan Unmarshal
func (infos *NotificationInfos) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &infos)
}

func (infos NotificationInfos) ToSlice() []*NotificationInfo {
	return infos
}
