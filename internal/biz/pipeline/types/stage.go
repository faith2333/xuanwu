package types

type Stage struct {
	Name         string                 `json:"name"`
	Code         string                 `json:"code"`
	ExecutorName string                 `json:"executorName"`
	ExecutorInfo map[string]interface{} `json:"executorInfo"`
	NextStages   []*Stage               `json:"nextStages"`
}
