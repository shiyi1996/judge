package models

import (
	"testing"
)

func TestPlayerCreate(t *testing.T) {
	InitAllInTest()

	player := Player{UserId: 1}
	_, err := Player{}.Create(&player)
	if err != nil {
		t.Error("Create() error")
	}
}
func TestPlayerRemove(t *testing.T) {
	InitAllInTest()

	err := Player{}.Remove(1)
	if err != nil {
		t.Error("Remove() error")
	}

}
func TestPlayerUpdate(t *testing.T) {
	InitAllInTest()

	player := Player{UserId: 1}

	err := Player{}.Update(&player)
	if err != nil {
		t.Error("Update() failed. Error:", err)
	}
}

func TestPlayerGetById(t *testing.T) {
	InitAllInTest()
	player := Player{UserId: 1}
	Player{}.Create(&player)

	getById, err := Player{}.GetById(player.Id)
	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getById != player {
		t.Error("GetById() failed. Error:", err)
	}
}
func TestPlayerGetByContestId(t *testing.T) {
	InitAllInTest()
	player := Player{ContestId: 1}
	Player{}.Create(&player)

	getByContestId, err := Player{}.GetByContestId(player.ContestId)
	if err != nil {
		t.Error("GetByContestId() failed.Error:", err)
	}

	if *getByContestId != player {
		t.Error("GetByContestId() :", err)
	}
}

func TestPlayerGetByUserId(t *testing.T) {
	InitAllInTest()
	player := Player{UserId: 6}
	Player{}.Create(&player)

	getById, err := Player{}.GetByUserId(player.UserId)

	if err != nil {
		t.Error("GetByUserId() failed.Error:", err)
	}

	if *getById != player {
		t.Error("GetByUserId() failed. Error:", err)
	}
}
