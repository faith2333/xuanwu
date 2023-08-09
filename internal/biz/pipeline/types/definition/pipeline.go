package definition

// Pipeline The definition of pipeline, which is formatted in YAML.
type Pipeline struct {
	// the global unique identifier for the pipeline
	Name            string      `yaml:"name"`
	GlobalVariables []*Variable `yaml:"globalVariables"`
	// two-dimensional array, the stages in first level is order relationship, it will be executed one by one,
	// the stages in second level is parallel relationship, will be executed concurrently.
	// all stages will parse as a DAG later.
	Stages [][]*Stage `yaml:"stages"`
}

func (pipe *Pipeline) Validate() error {
	if pipe.Name == "" {
		return ErrNameIsEmpty
	}

	nameExists := make(map[string]struct{})
	globalVars := make(map[string]*Variable)
	for _, v := range pipe.GlobalVariables {
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
