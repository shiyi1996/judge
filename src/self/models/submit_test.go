package models

import (
	"testing"
)

func TestSubmitCreate(t *testing.T) {
	InitAllInTest()

	submit := &Submit{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: "AC", ResultDes: "haha", CaseResult: "aaa", Code: "sskajka"}
	if _, err := submit.Create(submit); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestSubmitUpdate(t *testing.T) {
	InitAllInTest()

	submit := &Submit{Id: 1, ResultDes: "haha"}
	if err := submit.Update(submit); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestSubmitRemove(t *testing.T) {
	InitAllInTest()

	var submit Submit
	if err := submit.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestSubmitGetById(t *testing.T) {
	InitAllInTest()

	submit := &Submit{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: "AC", ResultDes: "haha", CaseResult: "aaa", Code: "sskajka"}
	submit.Create(submit)

	getSubmit, err := submit.GetById(submit.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getSubmit != *submit {
		t.Error("GetById() failed:", "%v != %v", submit, getSubmit)
	}
}
func TestSubmitQueryBySubmit(t *testing.T) {
	InitAllInTest()

	var submit Submit
	if _, err := submit.QueryBySubmit(2, 2, "2", "2"); err != nil {
		t.Error("QueryBySubmit() failed:", err)
	}
}
