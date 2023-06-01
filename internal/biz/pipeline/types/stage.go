package types

type Stage struct {
	Name         string                 `json:"name"`
	Code         string                 `json:"code"`
	ExecutorName string                 `json:"executorName"`
	ExecutorInfo map[string]interface{} `json:"executorInfo"`
	RunAfter     []*Stage               `json:"runAfter"`
}

func (s *Stage) NodeName() string {
	return s.Name
}

func (s *Stage) PrevNodeNames() []string {
	res := make([]string, 0)
	for _, prevStage := range s.RunAfter {
		res = append(res, prevStage.Name)
	}
	return res
}
