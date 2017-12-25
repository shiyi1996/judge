package models

import (
	"testing"
)

func TestLetterCreate(t *testing.T) {
	InitAllInTest()
	letter := Letter{IsRead: true}
	_, err := letter.Create(&letter)
	if err != nil {
		t.Error("Create() failed. Error:", err)
	}

}

func TestLetterRemove(t *testing.T) {
	InitAllInTest()
	err := Letter{}.Remove(1)
	if err != nil {
		t.Error("Remove() failed. Error:", err)
	}

}

func TestLetterUpdate(t *testing.T) {
	InitAllInTest()

	letter := &Letter{Content: "sasasaas", IsRead: false}
	Letter{}.Create(letter)

	err := Letter{}.Update(&Letter{Content: "dgdgdgd", IsRead: false})
	if err != nil {
		t.Error("Update() failed. Error:", err)
	}
}

func TestLetterGetById(t *testing.T) {
	InitAllInTest()
	letter := &Letter{Content: "sasasaas", IsRead: false}
	Letter{}.Create(letter)

	getById, err := Letter{}.GetById(letter.Id)

	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getById != *letter {
		t.Error("GetById() failed. Error:", getById, letter)
	}
}

func TestLetterQueryByIsRead(t *testing.T) {
	InitAllInTest()
	letter := &Letter{Content: "sasasaas", IsRead: false}

	queryByIsRead, err := Letter{}.QueryByIsRead(letter.IsRead)
	if err != nil && queryByIsRead[0] != *letter {
		t.Error("QueryByIsRead() failed. Error:", err)
	}

}
