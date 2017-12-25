package models

import "testing"

func TestProblemUserCreate(t *testing.T) {
	InitAllInTest()

	problemUser := &ProblemUser{Titile: "sadas", Description: "1111"}
	if _, err := problemUser.Create(problemUser); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestProblemUserRemove(t *testing.T) {
	InitAllInTest()

	var problemUser ProblemUser
	if err := problemUser.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestProblemUserUpdate(t *testing.T) {
	InitAllInTest()

	problemUser := &ProblemUser{Titile: "sadas", Description: "asdasdasd"}
	if err := problemUser.Update(problemUser); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestProblemUserGetById(t *testing.T) {
	InitAllInTest()

	problemUser := &ProblemUser{Titile: "sadas", Description: "fffff"}
	ProblemUser{}.Create(problemUser)

	getProblemUser, err := ProblemUser{}.GetById(problemUser.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblemUser != *problemUser {
		t.Error("GetById() failed:", "%v != %v", problemUser, getProblemUser)
	}
}

func TestProblemUserQueryByUserId(t *testing.T) {
	InitAllInTest()
	problemUser := &ProblemUser{Titile: "sadas", Description: "fffff", UserId: 1}
	ProblemUser{}.Create(problemUser)

	getProblemUser, err := ProblemUser{}.QueryByUserId(problemUser.UserId)

	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblemUser[0] != *problemUser {
		t.Error("GetById() failed:", "%v != %v", problemUser, getProblemUser)
	}

}
