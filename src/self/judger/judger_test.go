/**
 * Created by shiyi on 2017/12/21.
 * Email: shiyi@fightcoder.com
 */

package judger

import "testing"

func TestDoJudgeInDocker(t *testing.T) {
	judge := Judger{
		CodeFileName: "1.cpp",
		WorkDir:      "/Users/shiyi/project/fightcoder-judge/tmp/",
	}

	judge.doJudgeInDocker("default", "cpp")
}
