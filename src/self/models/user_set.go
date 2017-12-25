package models

import (
	. "self/commons/store"
)

type UserSet struct {
	Id        int64
	UserId    int64  //用户Id
	EditorSet string //编辑器设置(JSON)
}

//增加
func (this UserSet) Create(uset *UserSet) (int64, error) {
	_, err := OrmWeb.Insert(uset)
	if err != nil {
		return 0, err
	}
	return uset.Id, nil
}

//删除
func (this UserSet) Remove(id int64) error {
	userSet := UserSet{}
	_, err := OrmWeb.Id(id).Delete(userSet)
	return err
}

//修改
func (this UserSet) Update(userSet *UserSet) error {
	_, err := OrmWeb.AllCols().ID(userSet.Id).Update(userSet)
	return err
}

//查询
func (this UserSet) GetById(id int64) (*UserSet, error) {
	userSet := new(UserSet)
	has, err := OrmWeb.Id(id).Get(userSet)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userSet, nil
}

func (this UserSet) GetByUserId(userid int64) (*UserSet, error) {
	userset := new(UserSet)
	has, err := OrmWeb.Where("user_id = ?", userid).Get(userset)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userset, nil
}
