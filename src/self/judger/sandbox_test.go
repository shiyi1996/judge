package judger

import "testing"

var sandbox = Sandbox{
	JudgeType:   "default",
	Language:    "c++",
	TimeLimit:   1,
	MemoryLimit: 128000000,
	OutputLimit: 100,
}

func TestDoJudgeInDocker(t *testing.T) {
	sandbox.doJudgeInDocker("/Users/shiyi/project/fightcoder/judge/tmp")
}
