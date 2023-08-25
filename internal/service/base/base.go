package base

import (
	"encoding/json"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
)

type Base struct{}

func (b *Base) TransformJSON(source, target interface{}) error {
	sBytes, err := json.Marshal(source)
	if err != nil {
		return errors.Wrap(err, "marshal source failed")
	}

	return json.Unmarshal(sBytes, &target)
}

func (b *Base) PBStructUnmarshal(source *structpb.Struct, target interface{}) error {
	sBytes, err := json.Marshal(source.AsMap())
	if err != nil {
		return errors.Wrap(err, "marshal source failed")
	}

	return json.Unmarshal(sBytes, &target)
}

func (b *Base) NewPBStruct(source interface{}) (*structpb.Struct, error) {
	sBytes, err := json.Marshal(source)
	if err != nil {
		return nil, errors.Wrap(err, "marshal source failed")
	}

	sMap := make(map[string]interface{})
	err = json.Unmarshal(sBytes, &sMap)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal source failed")
	}

	return structpb.NewStruct(sMap)
}
