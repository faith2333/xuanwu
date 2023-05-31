package pipeline

import (
	"github.com/faith2333/xuanwu/internal/biz/pipeline/parseengine"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine"
)

type Pipeline struct {
	ProcessEngine processengine.Interface
	ParseEngine   parseengine.Interface
}
