package models

import (
	"testing"
)

func TestGroupCreate(t *testing.T) {

	InitAllInTest()
	group := Group{Name: "dsa"}
	_, err := group.Create(&group)
	if err != nil {
		t.Error("Create() failed. Error:", err)
	}

}

func TestGroup_Remove(t *testing.T) {

	InitAllInTest()
	group := new(Group)
	err := group.Remove(1)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestGroupUpdate(t *testing.T) {
	InitAllInTest()
	group := new(Group)
	group.Id = 2
	group.Description = "111111"
	err := group.Update(group)
	if err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestGroupGetById(t *testing.T) {
	InitAllInTest()
	group := Group{Description: "sda"}
	Group{}.Create(&group)

	getGroupById, err := Group{}.GetById(group.Id)

	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getGroupById != group {
		t.Error("GetById() failed. Error:", err)
	}

}

func TestGroupGetByName(t *testing.T) {
	InitAllInTest()
	group := Group{Description: "sda", Name: "cscas"}
	Group{}.Create(&group)

	getGroupByName, err := Group{}.GetByName(group.Name)

	if err != nil {
		t.Error("GetByName() failed:", err.Error())
	}

	if *getGroupByName != group {
		t.Error("GetByName() failed. Error:", err)
	}
}
