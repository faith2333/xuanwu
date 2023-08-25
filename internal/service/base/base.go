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

func (b *Base) SlicePBStructUnmarshal(sources []*structpb.Struct, target interface{}) error {
	s := make([]map[string]interface{}, 0)
	for _, source := range sources {
		s = append(s, source.AsMap())
	}

	sBytes, err := json.Marshal(s)
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

func (b *Base) NewPBStructSlice(source interface{}) ([]*structpb.Struct, error) {
	sBytes, err := json.Marshal(source)
	if err != nil {
		return nil, errors.Wrap(err, "marshal source failed")
	}

	sMap := make([]map[string]interface{}, 0)
	err = json.Unmarshal(sBytes, &sMap)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal source failed")
	}

	s := make([]*structpb.Struct, 0)
	for _, m := range sMap {
		pbStruct, err := structpb.NewStruct(m)
		if err != nil {
			return nil, errors.Wrap(err, "new pb struct failed")
		}
		s = append(s, pbStruct)
	}

	return s, nil
}
