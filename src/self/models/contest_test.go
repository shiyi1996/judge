package models

import (
	"testing"
)

func TestContestCreate(t *testing.T) {

	InitAllInTest()
	contest := Contest{}
	_, err := contest.Create(&contest)
	if err != nil {
		t.Error("Error: ", err)
	}

}

func TestContestRemove(t *testing.T) {
	InitAllInTest()
	contest := new(Contest)
	err := contest.Remove(1)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}

}
func TestContestUpdate(t *testing.T) {

	InitAllInTest()
	contest := new(Contest)
	contest.Id = 2
	contest.UserId = 1
	contest.Description = "ddddd"
	err := contest.Update(contest)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestContestGetById(t *testing.T) {

	InitAllInTest()
	contest := Contest{Description: "wweweq"}
	Contest{}.Create(&contest)

	getById, err := contest.GetById(contest.Id)
	if err != nil {
		t.Error("GetById() failed. Error:", err)
	}

	if *getById != contest {
		t.Error("GetById() failed. Error:", err)
	}
}

func TestContestGetByUserId(t *testing.T) {

	InitAllInTest()
	contest := &Contest{Description: "1", UserId: 2}
	Contest{}.Create(contest)

	getByUserId, err := Contest{}.GetByUserId(2)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}

	if *getByUserId != *contest {
		t.Error("GetByUserId() failed:", contest, "!=", getByUserId)
	}
}

func TestContestQueryByType(t *testing.T) {

	InitAllInTest()
	contest := new(Contest)

	_, err := contest.QueryByType(0)
	if err != nil {
		t.Error("QueryByType() failed. Error:", err)
	}

}
