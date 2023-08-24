package parseengine

import (
	"context"
	"fmt"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"testing"
)

var TestPipeline = &definition.Pipeline{
	Name: "test",
	Code: "test",
	Type: definition.PipelineTypeCICD,
	GlobalVariables: []*definition.Variable{
		{
			Key:      "env",
			Type:     definition.VariableTypeSlice,
			Required: true,
		},
	},
	Stages: [][]*definition.Stage{
		{
			{
				Name:         "build",
				ExecutorName: "POD",
				ExecuteInfo: map[string]interface{}{
					"test": "${test}$",
				},
				Repeat:             true,
				RepeatFromVariable: "env",
				NextStages:         nil,
			},
		},
		{
			{
				Name:         "deploy",
				ExecutorName: "POD",
				ExecuteInfo: map[string]interface{}{
					"testKey": "testValue",
				},
			},
		},
	},
}

func TestDefaultEngine_ParseAndGenerate(t *testing.T) {
	eng := &defaultEngine{}

	_, nodes, err := eng.ParseAndGenerate(context.Background(), TestPipeline, map[string]interface{}{"env": []string{"test", "test2"}})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v \n", nodes)

}
