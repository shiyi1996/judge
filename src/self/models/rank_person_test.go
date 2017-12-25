package models

import (
	"testing"
)

func TestRankPersonCreate(t *testing.T) {
	InitAllInTest()

	rankPerson := &RankPerson{UserId: 3, AcNumber: 1, Grade: 2}
	if _, err := rankPerson.Add(rankPerson); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestRankPersonUpdate(t *testing.T) {
	InitAllInTest()

	rankPerson := &RankPerson{Id: 1, UserId: 2, AcNumber: 2, Grade: 2}
	if err := rankPerson.Update(rankPerson); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestRankPersonRemove(t *testing.T) {
	InitAllInTest()

	var rankPerson RankPerson
	if err := rankPerson.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestRankPersonGetById(t *testing.T) {
	InitAllInTest()

	rankPerson := &RankPerson{UserId: 3, AcNumber: 1, Grade: 2}
	rankPerson.Add(rankPerson)

	getRankPerson, err := rankPerson.GetById(rankPerson.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getRankPerson != *rankPerson {
		t.Error("GetById() failed:", "%v != %v", rankPerson, getRankPerson)
	}
}
func TestRankPersonQueryByRankPerson(t *testing.T) {
	InitAllInTest()

	rankPerson := &RankPerson{UserId: 2}
	if _, err := rankPerson.QueryByRankPerson(rankPerson); err != nil {
		t.Error("QueryByRankPerson() failed:", err)
	}
}
