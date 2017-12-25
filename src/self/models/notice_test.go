package models

import (
	"testing"
)

func TestNoticeCreate(t *testing.T) {

	InitAllInTest()
	notice := Notice{UserId: 1}
	_, err := Notice{}.Create(&notice)
	if err != nil {
		t.Error("Create() failed. Error:", err)
	}
}

func TestNoticeRemove(t *testing.T) {
	InitAllInTest()

	err := Notice{}.Remove(1)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestNoticeUpdate(t *testing.T) {
	InitAllInTest()
	notice := Notice{IsRead: false, UserId: 1}
	err := Notice{}.Update(&notice)
	if err != nil {
		t.Error("Update() failed. Error:", err)
	}
}

func TestNoticeGetById(t *testing.T) {

	InitAllInTest()
	notice := Notice{UserId: 1}
	Notice{}.Create(&notice)

	getById, err := Notice{}.GetById(notice.Id)
	if err != nil {
		t.Error("GetById() failed. Error:", err)
	}
	if *getById != notice {
		t.Error("GetById() failed. Error:", err)
	}
}

func TestNoticeGetByUserId(t *testing.T) {
	InitAllInTest()
	notice := Notice{UserId: 2}
	Notice{}.Create(&notice)

	getById, err := Notice{}.GetByUserId(notice.UserId)

	if err != nil {
		t.Error("GetByUserId() failed. Error:", err)
	}
	if *getById != notice {
		t.Error("GetByUserId() failed. Error:", getById, notice)
	}
}
func TestNotice_QueryByIsRead(t *testing.T) {

	InitAllInTest()
	notice := Notice{UserId: 3, IsRead: true}
	Notice{}.Create(&notice)

	getById, err := Notice{}.QueryByIsRead(notice.IsRead)

	if err != nil {
		t.Error("QueryByIsRead() failed. Error:", err)
	}
	if getById[0] != notice {
		t.Error("QueryByIsRead() failed. Error:", err)
	}
}
