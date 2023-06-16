package base

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Base struct{}

func (b *Base) TransformJSON(source, target interface{}) error {
	sBytes, err := json.Marshal(source)
	if err != nil {
		return errors.Wrap(err, "marshal source failed")
	}

	return json.Unmarshal(sBytes, &target)
}
