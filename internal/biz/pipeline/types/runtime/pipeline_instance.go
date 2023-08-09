package runtime

import "time"

type Instance struct {
	// the global unique identifier of the pipeline instance
	Code string `json:"code"`
	// the display name of the pipeline instance, it's inherited from pipeline definition
	Name string `json:"name"`
	// the code of the pipeline definition
	DefinitionCode  string                 `json:"definitionCode"`
	GlobalVariables map[string]interface{} `json:"globalVariables"`
	Status          Status                 `json:"status"`
	StartTime       time.Time              `json:"startTime"`
	FinishTime      time.Time              `json:"finishTime"`
	TimeCostSec     int64                  `json:"timeCostSec"`
}
