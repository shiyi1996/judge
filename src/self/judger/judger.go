/**
 * Created by shiyi on 2017/10/16.
 * Email: shiyi@fightcoder.com
 */

package judger

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"self/commons/components"
	"self/commons/g"
	"self/models"

	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
)

var codeSuffixMap = map[string]string{
	"c":      ".c",
	"c++":    ".cpp",
	"java":   ".java",
	"golang": ".go",
	"python": ".py",
}

type Judger struct {
	SubmitType string         `json:"submit_type"` //提交类型
	SubmitId   int64          `json:"submit_id"`   //提交id
	Problem    models.Problem `json:"problem"`     //题目信息
	Submit     models.Submit  `json:"submit"`      //提交信息
	WorkDir    string
}

// TODO Java
func (this *Judger) DoJudge() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Judge Error : %v", err)
		}
	}()

	this.getSubmitData()
	this.getProblemData(this.Submit.ProblemType)
	this.createWorkDir()

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

	code, _ := sandbox.doJudgeInDocker(this.WorkDir)
	if code != 0 {
		this.getSubmitData()
		if this.Submit.Result == Normal ||
			this.Submit.Result == Waiting ||
			this.Submit.Result == Compiling ||
			this.Submit.Result == Running {
			this.Submit.Result = SystemError
			this.saveSubmit()
		}
	}
}

func (this *Judger) doSpecialJudge() {

}

func (this *Judger) doContestJudge() {

}

func (this *Judger) doTestJudge() {

}

func (this *Judger) createWorkDir() {
	cfg := g.Conf()
	this.WorkDir = getCurrentPath() + "/" + cfg.Judge.WorkDir + "/" +
		strconv.FormatInt(this.Submit.UserId, 10) + "_" +
		this.SubmitType + "_" + strconv.FormatInt(this.Submit.Id, 10)

	err := os.MkdirAll(this.WorkDir, 0777)
	if err != nil {
		panic("createWorkDir: " + err.Error())
	}
}

func (this *Judger) removeWorkDir() {
}

func (this *Judger) getCase() {
	minioCli := components.NewMinioCli()
	minioCli.DownloadCase(this.Problem.CaseData, this.WorkDir+"/case.zip")

	err := archiver.Zip.Open(this.WorkDir+"/case.zip", this.WorkDir+"/case")
	if err != nil {
		panic("getCase: " + err.Error())
	}
}

func (this *Judger) getCode() {
	minioCli := components.NewMinioCli()

	path := this.WorkDir + "/code" + codeSuffixMap[this.Submit.Language]
	minioCli.DownloadCode(this.Submit.Code, path)
}

func (this *Judger) saveSubmit() {
	switch this.SubmitType {
	case "submit":
		submit, err := models.Submit{}.GetById(this.SubmitId)
		if err != nil {
			panic(err)
		}
		submit.Result = this.Submit.Result
		submit.ResultDes = this.Submit.ResultDes
		submit.RunningTime = this.Submit.RunningTime
		submit.RunningMemory = this.Submit.RunningMemory

		err = models.Submit{}.Update(submit)
		if err != nil {
			panic(err)
		}

		log.Infof("saveSubmit: %#v", submit)

	case "submit_contest":
	case "submit_user":
	case "submit_test":
	}
}

func (this *Judger) getProblemData(problemType string) {
	var problemJson []byte

	switch problemType {
	case "real":
		{
			problem, err := models.Problem{}.GetById(this.Submit.ProblemId)
			if err != nil {
				panic("getProblemData-Problem: " + err.Error())
			}
			problemJson, err = json.Marshal(problem)
			break
		}
	case "user":
		{
			problemUser, err := models.ProblemUser{}.GetById(this.Submit.ProblemId)
			if err != nil {
				panic("getProblemData-ProblemUser: " + err.Error())
			}
			problemJson, err = json.Marshal(problemUser)
			break
		}
	default:
		panic("getProblemData: not recognized ProblemType " + this.Submit.ProblemType)
	}

	if err := json.Unmarshal(problemJson, &this.Problem); err != nil {
		panic("getProblemData: " + err.Error())
	}

	log.Infof("getProblemData: %#v\n", this.Problem)
}

func (this *Judger) getSubmitData() {
	var submitJson []byte

	switch this.SubmitType {
	case "submit":
		{
			submit, err := models.Submit{}.GetById(this.SubmitId)
			if err != nil {
				panic("getSubmitData-Submit: " + err.Error())
			}
			submitJson, err = json.Marshal(submit)
			break
		}
	case "submit_user":
		{
			submitUser, err := models.SubmitUser{}.GetById(this.SubmitId)
			if err != nil {
				panic("getSubmitData-SubmitUser: " + err.Error())
			}
			submitJson, err = json.Marshal(submitUser)
			break
		}
	case "submit_contest":
		{
			submitContest, err := models.SubmitContest{}.GetById(this.SubmitId)
			if err != nil {
				panic("getSubmitData-SubmitContest: " + err.Error())
			}
			submitJson, err = json.Marshal(submitContest)
			break
		}
	case "submit_test":
		{

		}
	default:
		panic("getSubmitData: not recognized submitType " + this.SubmitType)
	}

	if err := json.Unmarshal(submitJson, &this.Submit); err != nil {
		panic("getSubmitData: " + err.Error())
	}

	log.Infof("getSubmitData: %#v\n", this.Submit)
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("getCurrentPath: " + err.Error())
	}
	return dir
}
