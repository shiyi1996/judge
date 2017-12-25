package models

import (
	"testing"
)

func TestMemberCreate(t *testing.T) {
	InitAllInTest()
	member := Member{UserId: 1}
	_, err := Member{}.Create(&member)
	if err != nil {
		t.Error("Create() failed. Error:", err)
	}
}

func TestMemberUpdate(t *testing.T) {

	InitAllInTest()
	member := Member{UserId: 1}
	Member{}.Create(&member)
	err := Member{}.Update(&Member{UserId: 1})
	if err != nil {
		t.Error("Create() failed. Error:", err)
	}
}

func TestMemberGetById(t *testing.T) {
	InitAllInTest()
	member := Member{UserId: 1}
	getById, err := Member{}.GetById(member.Id)
	if err != nil && getById != &member {
		t.Error("GetById() failed. Error:", err)
	}
}

func TestMemberRemove(t *testing.T) {

	InitAllInTest()

	err := Member{}.Remove(1)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestMember_GetByGroupId(t *testing.T) {
	InitAllInTest()
	member := new(Member)
	_, err := member.GetByGroupId(1)
	if err != nil {
		t.Error("GetByGroupId() failed. Error:", err)
	}
}

func TestMember_GetByUserId(t *testing.T) {
	InitAllInTest()

	_, err := Member{}.GetByUserId(1)
	if err != nil {
		t.Error("GetByUserId() failed. Error:", err)
	}
}

func TestMember_QueryByRole(t *testing.T) {
	InitAllInTest()
	_, err := Member{}.QueryByRole("")

	if err != nil {
		t.Error("QueryByRole() failed. Error:", err)
	}
}
