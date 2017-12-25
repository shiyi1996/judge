package models

import (
	"testing"
)

func TestRankGroupCreate(t *testing.T) {
	InitAllInTest()

	rankGroup := &RankGroup{GroupId: 3, AcNumber: 1, Grade: 2}
	if _, err := rankGroup.Add(rankGroup); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestRankGroupUpdate(t *testing.T) {
	InitAllInTest()

	rankGroup := &RankGroup{Id: 1, GroupId: 2, AcNumber: 2, Grade: 2}
	if err := rankGroup.Update(rankGroup); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestRankGroupRemove(t *testing.T) {
	InitAllInTest()

	var rankGroup RankGroup
	if err := rankGroup.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestRankGroupGetById(t *testing.T) {
	InitAllInTest()

	rankGroup := &RankGroup{GroupId: 3, AcNumber: 1, Grade: 2}
	rankGroup.Add(rankGroup)

	getRankGroup, err := rankGroup.GetById(rankGroup.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getRankGroup != *rankGroup {
		t.Error("GetById() failed:", "%v != %v", rankGroup, getRankGroup)
	}
}
func TestRankGroupQueryByRankGroup(t *testing.T) {
	InitAllInTest()

	RankGroup := &RankGroup{GroupId: 2}
	if _, err := RankGroup.QueryByRankGroup(RankGroup); err != nil {
		t.Error("QueryBySubmit() failed:", err)
	}
}
