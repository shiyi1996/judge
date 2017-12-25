package models

import (
	"testing"
)

func TestUserSetCreate(t *testing.T) {
	InitAllInTest()

	userSet := &UserSet{UserId: 3, EditorSet: "aaaaa"}
	if _, err := userSet.Create(userSet); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}

func TestUserSetUpdate(t *testing.T) {
	InitAllInTest()

	userSet := &UserSet{Id: 1, UserId: 4, EditorSet: "哈哈哈"}
	if err := userSet.Update(userSet); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}

func TestUserSetRemove(t *testing.T) {
	InitAllInTest()

	var userSet UserSet
	if err := userSet.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestUserSetGetById(t *testing.T) {
	InitAllInTest()

	userSet := &UserSet{UserId: 3, EditorSet: "aaaaa"}
	userSet.Create(userSet)
	userSetGet, err := userSet.GetById(userSet.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *userSetGet != *userSet {
		t.Error("GetById() failed:", "%v != %v", userSetGet, userSet)
	}

}
func TestUserSetGetByUserId(t *testing.T) {
	InitAllInTest()

	userSet := &UserSet{UserId: 20, EditorSet: "aaaaa"}
	userSet.Create(userSet)
	userSetGet, err := userSet.GetByUserId(userSet.UserId)
	if err != nil {
		t.Error("GetByUserId() failed:", err.Error())
	}

	if *userSetGet != *userSet {
		t.Error("GetByUserId() failed:", "%v != %v", userSetGet, userSet)
	}
}
