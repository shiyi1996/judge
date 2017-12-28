package judger

import (
	"encoding/json"
	"fmt"
)

type Sandbox struct {
	SubmitType  string
	SubmitId    int64
	JudgeType   string
	Language    string
	TimeLimit   int64
	MemoryLimit int64
	OutputLimit int64
}

func (this *Sandbox) doJudgeInDocker(workDir string) {
	dockerCli := NewDockerCli()

	jsonData, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	cmd := []string{"/fightcoder-sandbox/sandbox",
		"-judge_data", string(jsonData),
	}

	fmt.Println("CMD: ", cmd)

	dockerCli.RunContainer("sandbox", cmd, workDir)
}
