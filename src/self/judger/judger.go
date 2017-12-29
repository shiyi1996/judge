/**
 * Created by shiyi on 2017/10/16.
 * Email: shiyi@fightcoder.com
 */

package judger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"self/commons/components"
	"self/models"

	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
)

var codeSuffixMap = map[string]string{
	"c":      ".cpp",
	"c++":    ".cpp",
	"java":   ".java",
	"golang": ".go",
	"python": ".py",
}

type Judger struct {
	SubmitType      string         `json:"submit_type"`       //提交类型
	SubmitId        int64          `json:"submit_id"`         //提交id
	ProblemBankType string         `json:"problem_bank_type"` //题库类型
	ProblemId       int64          `json:"problem_id"`        //题目Id
	Problem         models.Problem `json:"problem"`           //题目信息
	Submit          models.Submit  `json:"submit"`            //提交信息
	WorkDir         string
}

func (this *Judger) DoJudge() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Judge Error : %v", err)
		}
	}()

	this.getProblemData()
	this.getSubmitData()
	this.createWorkDir()

	//this.getCases()
	//this.getCode()

	this.doJudge()
	//if this.Problem.IsSpecialJudge {
	//	this.doSpecialJudge()
	//} else {
	//	this.doJudge()
	//}

	this.removeWorkDir()
}

func (this *Judger) doJudge() {
	this.getCode()
	this.getCase()

	//如果判题类型没问题，获取特判代码，不必分开
	sandbox := Sandbox{
		SubmitType:  this.SubmitType,
		SubmitId:    this.SubmitId,
		JudgeType:   "default",
		Language:    this.Submit.Language,
		TimeLimit:   int64(this.Problem.TimeLimit),
		MemoryLimit: int64(this.Problem.MemoryLimit),
		OutputLimit: 100,
	}

	sandbox.doJudgeInDocker(this.WorkDir)
}

func (this *Judger) doSpecialJudge() {

}

func (this *Judger) doContestJudge() {

}

func (this *Judger) doTestJudge() {

}

func (this *Judger) createWorkDir() {
	this.WorkDir = getCurrentPath() + "/tmp/" + strconv.FormatInt(this.Submit.UserId, 10) + "_" + this.SubmitType + "_" + strconv.FormatInt(this.Submit.Id, 10)

	err := os.MkdirAll(this.WorkDir, 0777)
	if err != nil {
		panic(err)
	}
}

func (this *Judger) removeWorkDir() {
}

func (this *Judger) getCase() {
	minioCli := components.NewMinioCli()

	minioCli.DownloadCase(this.Problem.Case, this.WorkDir+"/case.zip")

	err := archiver.Zip.Open(this.WorkDir+"/case.zip", this.WorkDir)
	if err != nil {
		panic(err)
	}
}

func (this *Judger) getCode() {
	minioCli := components.NewMinioCli()

	path := this.WorkDir + "/code" + codeSuffixMap[this.Submit.Language]
	minioCli.DownloadCode(this.Submit.Code, path)
}

func (this *Judger) getProblemData() {
	var problemJson []byte

	switch this.ProblemBankType {
	case "problem":
		{
			problem, err := models.Problem{}.GetById(this.ProblemId)
			if err != nil {
				panic(err)
			}
			problemJson, err = json.Marshal(problem)
			break
		}
	case "problem_user":
		{
			problemUser, err := models.ProblemUser{}.GetById(this.ProblemId)
			if err != nil {
				panic(err)
			}
			problemJson, err = json.Marshal(problemUser)
			break
		}
	default:
		panic("not recognized problemBankType " + this.ProblemBankType)
	}

	if err := json.Unmarshal(problemJson, &this.Problem); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", this.Problem)
}

func (this *Judger) getSubmitData() {
	var submitJson []byte

	switch this.SubmitType {
	case "submit":
		{
			submit, err := models.Submit{}.GetById(this.SubmitId)
			if err != nil {
				panic(err)
			}
			submitJson, err = json.Marshal(submit)
			break
		}
	case "submit_user":
		{
			submitUser, err := models.SubmitUser{}.GetById(this.SubmitId)
			if err != nil {
				panic(err)
			}
			submitJson, err = json.Marshal(submitUser)
			break
		}
	case "submit_contest":
		{
			submitContest, err := models.SubmitContest{}.GetById(this.SubmitId)
			if err != nil {
				panic(err)
			}
			submitJson, err = json.Marshal(submitContest)
			break
		}
	case "submit_test":
		{

		}
	default:
		panic("not recognized submitType " + this.SubmitType)
	}

	if err := json.Unmarshal(submitJson, &this.Submit); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", this.Submit)
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
