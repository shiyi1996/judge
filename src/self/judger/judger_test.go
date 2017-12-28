/**
 * Created by shiyi on 2017/12/21.
 * Email: shiyi@fightcoder.com
 */

package judger

import (
	"testing"

	"self/models"
)

var judger = Judger{
	Submit: models.Submit{
		Code:     "2.cpp",
		Language: "c++",
	},
	Problem: models.Problem{
		TimeLimit:   1,
		MemoryLimit: 128000,
		InputCase:   "case.zip",
	},
	WorkDir: "/Users/shiyi/project/fightcoder/judge/tmp",
}

func TestDoJudge(t *testing.T) {
	models.InitAllInTest()

	judger.doJudge()
}

func TestGetCode(t *testing.T) {
	models.InitAllInTest()

	judger.getCode()
}

func TestGetCase(t *testing.T) {
	models.InitAllInTest()

	judger.getCase()
}
