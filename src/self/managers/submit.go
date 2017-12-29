package managers

import (
	"self/judger"
	"self/models"
)

type SubmitManager struct {
}

func (this SubmitManager) ChangeSubmitResult(submitType string, submitId int64, result judger.Result) error {
	switch submitType {
	case "submit":
		submit := models.Submit{
			Id: submitId,
			//Result:result.ResultCode,
			ResultDes:     result.ResultDes,
			RunningTime:   result.RunningTime,
			RunningMemory: result.RunningMemory,
		}
		err := models.Submit{}.Update(&submit)
		if err != nil {
			panic(err)
		}
	case "submit_contest":
	case "submit_user":
	case "submit_test":
	}

	return nil
}