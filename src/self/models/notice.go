package models

import (
	. "self/commons/store"
)

type Notice struct {
	Id      int64  //通知id
	UserId  int64  //用户Id
	Time    int64  //时间
	Content string //内容
	IsRead  bool   //是否已读
}

//增加
func (this Notice) Create(notice *Notice) (int64, error) {
	_, err := OrmWeb.Insert(notice)
	if err != nil {
		return 0, err
	}
	return notice.Id, nil
}

//删除
func (this Notice) Remove(id int64) error {
	notice := Notice{}
	_, err := OrmWeb.Id(id).Delete(notice)
	return err
}

//修改
func (this Notice) Update(notice *Notice) error {
	_, err := OrmWeb.Id(notice.Id).Update(notice)
	return err
}

//查询
func (this Notice) GetById(id int64) (*Notice, error) {
	notice := new(Notice)
	has, err := OrmWeb.Id(id).Get(notice)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return notice, nil
}

func (this Notice) GetByUserId(userId int64) (*Notice, error) {
	notice := new(Notice)
	has, err := OrmWeb.Where("user_id = ?", userId).Get(notice)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return notice, nil
}

func (this Notice) QueryByIsRead(read bool) ([]Notice, error) {
	noticeList := make([]Notice, 0)
	err := OrmWeb.Where("is_read = ?", read).Find(&noticeList)
	if err != nil {
		return nil, err
	}
	return noticeList, nil
}
