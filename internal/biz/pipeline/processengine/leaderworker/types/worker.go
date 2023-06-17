package types

type WorkerID string

func (id WorkerID) String() string {
	return string(id)
}
