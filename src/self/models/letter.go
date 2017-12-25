package models

import (
	. "self/commons/store"
)

type Letter struct {
	Id            int64  //站内信id
	Content       string //内容
	SendUserId    int64  //发送者id
	ReceiveUserId int64  //接收者id
	SendTime      int64  //发送时间
	IsRead        bool   //是否已读
}

//增加
func (this Letter) Create(letter *Letter) (int64, error) {
	_, err := OrmWeb.Insert(letter)
	if err != nil {
		return 0, err
	}
	return letter.Id, nil
}

//删除
func (this Letter) Remove(id int64) error {
	letter := Letter{}
	_, err := OrmWeb.Id(id).Delete(letter)
	return err
}

//修改
func (this Letter) Update(letter *Letter) error {
	_, err := OrmWeb.Id(letter.Id).Update(letter)
	return err
}

//查询
func (this Letter) GetById(id int64) (*Letter, error) {
	letter := new(Letter)
	has, err := OrmWeb.Id(id).Get(letter)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return letter, nil
}

func (this Letter) QueryByIsRead(read bool) ([]Letter, error) {
	letterList := make([]Letter, 0)
	err := OrmWeb.Where("is_read = ?", read).Find(&letterList)
	if err != nil {
		return nil, err
	}
	return letterList, nil
}
