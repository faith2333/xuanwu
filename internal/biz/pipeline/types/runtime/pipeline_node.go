package runtime

import "time"

type Node struct {
	Name string `json:"name"`
	Code string `json:"code"`
	// the code of the pipeline instance
	InstanceCode string                 `json:"instanceCode"`
	ExecutorName string                 `json:"executorName"`
	ExecuteInfo  map[string]interface{} `json:"executeInfo"`
	// the front nodes name of the node
	RunAfter    []string  `json:"runAfter"`
	StartTime   time.Time `json:"startTime"`
	FinishTime  time.Time `json:"finishTime"`
	TimeCostSec int64     `json:"timeCostSec"`
}

func (node *Node) NodeName() string {
	return node.Name
}

func (node *Node) PrevNodeNames() []string {
	return node.RunAfter
}
