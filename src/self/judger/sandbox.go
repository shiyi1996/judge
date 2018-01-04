package judger

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
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

	cmd := []string{"/sandbox/sandbox",
		"-judge_data", string(jsonData),
	}

	log.Infof("judge in docker sandbox; \ndata: %s \ncmd: %#v; workDir: %s",
		string(jsonData), cmd, workDir)

	dockerCli.RunContainer("sandbox", cmd, workDir)
}
