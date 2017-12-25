package models

import (
	"testing"
)

func TestUserRelationCreate(t *testing.T) {
	InitAllInTest()

	userRelation := &UserRelation{LeaderId: 1, FollowerId: 3}
	if _, err := userRelation.Create(userRelation); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestUserRelationUpdate(t *testing.T) {
	InitAllInTest()

	userRelation := &UserRelation{Id: 1, LeaderId: 1, FollowerId: 3}
	if err := userRelation.Update(userRelation); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserRelationRemove(t *testing.T) {
	InitAllInTest()

	var userRelation UserRelation
	if err := userRelation.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserRelationGetById(t *testing.T) {
	InitAllInTest()

	userRelation := &UserRelation{LeaderId: 1, FollowerId: 3}
	userRelation.Create(userRelation)

	getUserRelation, err := userRelation.GetById(userRelation.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUserRelation != *userRelation {
		t.Error("GetById() failed:", "%v != %v", userRelation, getUserRelation)
	}
}
func TestUserRelationQueryBySubmit(t *testing.T) {
	InitAllInTest()

	userRelation := &UserRelation{LeaderId: 2}
	if _, err := userRelation.QueryByUserRelation(userRelation); err != nil {
		t.Error("QueryByUserRelation() failed:", err)
	}
}
