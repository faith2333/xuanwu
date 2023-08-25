package definition

type Stage struct {
	// the unique identifier for the stage in the pipeline
	Name string `yaml:"name" json:"name"`
	// the executor name for the stage
	ExecutorName string                 `yaml:"executorName" json:"executorName"`
	ExecuteInfo  map[string]interface{} `yaml:"executeInfo" json:"executeInfo"`
	// if your specify Repeat as true, then you must specify the RepeatFromVariable.
	// and this stage will be repeated generated.
	Repeat bool `yaml:"repeat" json:"repeat"`
	// the Key of GlobalVariables in pipeline definition. and the type of this variable must be SLICE or MAP.
	RepeatFromVariable string `yaml:"repeatFromVariable" json:"repeatFromVariable"`
	// two-dimensional array, the stages in first level is order relationship, it will be executed one by one,
	// the stages in second level is parallel relationship, will be executed concurrently.
	NextStages [][]*Stage `yaml:"nextStages" json:"nextStages"`
}

func (stage *Stage) Validate(nameExists map[string]struct{}, globalVars map[string]*Variable) error {
	if stage.Name == "" {
		return ErrNameIsEmpty
	}

	if _, ok := nameExists[stage.Name]; ok {
		return ErrMultipleName
	}
	nameExists[stage.Name] = struct{}{}

	if stage.Repeat && stage.RepeatFromVariable == "" {
		return ErrRepeatFromVariableNotSpecified
	}

	if stage.Repeat {
		v, ok := globalVars[stage.RepeatFromVariable]
		if !ok {
			return ErrRepeatFromVariableNotFound
		}

		if !v.Type.IsMap() && !v.Type.IsSlice() {
			return ErrRepeatFromVariableTypeNotSupported
		}
	}

	for _, stages := range stage.NextStages {
		for _, s := range stages {
			err := s.Validate(nameExists, globalVars)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
