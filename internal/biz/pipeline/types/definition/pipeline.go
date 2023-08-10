package definition

// Pipeline The definition of pipeline, which is formatted in YAML.
type Pipeline struct {
	// the display name of the pipeline definition
	Name string `yaml:"name"`
	// the global unique identifier of the pipeline definition
	Code            string       `json:"code"`
	Type            PipelineType `yaml:"type"`
	GlobalVariables []*Variable  `yaml:"globalVariables"`
	// two-dimensional array, the stages in first level is order relationship, it will be executed one by one,
	// the stages in second level is parallel relationship, will be executed concurrently.
	// all stages will parse as a DAG later.
	Stages [][]*Stage `yaml:"stages"`
}

// Validate confirm that the definition of pipeline is legal.
func (pipe *Pipeline) Validate() error {
	if pipe.Name == "" {
		return ErrNameIsEmpty
	}

	if pipe.Type == "" {
		return ErrTypeIsEmpty
	}

	if !pipe.Type.IsSupported() {
		return ErrTypeIsNotSupport
	}

	nameExists := make(map[string]struct{})
	globalVars := make(map[string]*Variable)
	for _, v := range pipe.GlobalVariables {
		if !v.Type.IsSupported() {
			return ErrVariableNotSupported
		}
		globalVars[v.Key] = v
	}

	for _, stages := range pipe.Stages {
		for _, stage := range stages {
			err := stage.Validate(nameExists, globalVars)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
