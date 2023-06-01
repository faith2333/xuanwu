package types

type LogicTaskID string

func (tID LogicTaskID) String() string {
	return string(tID)
}

type LogicTask interface {
	GetID() LogicTaskID
	GetData() []byte
}

type defaultLogicTask struct {
	taskID LogicTaskID
	data   []byte
}

func NewDefaultLogicTask(id LogicTaskID, data []byte) LogicTask {
	return &defaultLogicTask{
		taskID: id,
		data:   data,
	}
}

func (task *defaultLogicTask) GetID() LogicTaskID {
	return task.taskID
}

func (task *defaultLogicTask) GetData() []byte {
	return task.data
}
