package runtime

import "strings"

type Status string

const (
	StatusInitialization Status = "INITIALIZATION"
	StatusRunning        Status = "RUNNING"
	StatusSuccess        Status = "SUCCESS"
	StatusFailed         Status = "FAILED"
	StatusCanceled       Status = "CANCELED"
)

var AllStatus = []Status{
	StatusInitialization,
	StatusRunning,
	StatusSuccess,
	StatusFailed,
	StatusCanceled,
}

func (s Status) String() string {
	return string(s)
}

func (s Status) Upper() Status {
	return Status(strings.ToUpper(s.String()))
}

func (s Status) IsSupported() bool {
	for _, aStatus := range AllStatus {
		if s.Upper() == aStatus {
			return true
		}
	}

	return false
}

func (s Status) IsInitialization() bool {
	return s.Upper() == StatusInitialization
}

func (s Status) Running() bool {
	return s.Upper() == StatusRunning
}

func (s Status) IsSuccess() bool {
	return s.Upper() == StatusSuccess
}

func (s Status) IsFailed() bool {
	return s.Upper() == StatusFailed
}

func (s Status) IsCanceled() bool {
	return s.Upper() == StatusRunning
}
