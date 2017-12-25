package models

import (
	"testing"
)

func TestSubmitUserCreate(t *testing.T) {
	InitAllInTest()

	submitUser := &SubmitUser{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: "AC", ResultDes: "haha", CaseResult: "aaa", Code: "sskajka"}
	if _, err := submitUser.Create(submitUser); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestSubmitUserUpdate(t *testing.T) {
	InitAllInTest()

	submitUser := &SubmitUser{Id: 1, ResultDes: "haha"}
	if err := submitUser.Update(submitUser); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestSubmitUserRemove(t *testing.T) {
	InitAllInTest()

	var submitUser SubmitUser
	if err := submitUser.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestSubmitUserGetById(t *testing.T) {
	InitAllInTest()

	submitUser := &SubmitUser{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: "AC", ResultDes: "haha", CaseResult: "aaa", Code: "sskajka"}
	submitUser.Create(submitUser)

	getSubmitUser, err := submitUser.GetById(submitUser.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getSubmitUser != *submitUser {
		t.Error("GetById() failed:", "%v != %v", submitUser, getSubmitUser)
	}
}
func TestSubmitUserQueryBySubmit(t *testing.T) {
	InitAllInTest()

	submitUser := &SubmitUser{ProblemId: 2}
	if _, err := submitUser.QueryByContestSubmit(submitUser); err != nil {
		t.Error("QueryBySubmit() failed:", err)
	}
}
