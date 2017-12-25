package models

import (
	. "self/commons/store"
)

type UserRelation struct {
	Id         int64
	LeaderId   int64 //领导者：即被关注者
	FollowerId int64 //追随者：即关注者
}

//增加
func (this UserRelation) Create(userRelation *UserRelation) (int64, error) {
	_, err := OrmWeb.Insert(userRelation)
	if err != nil {
		return 0, err
	}
	return userRelation.Id, nil
}

//删除
func (this UserRelation) Remove(id int64) error {
	userRelation := UserRelation{}
	_, err := OrmWeb.Id(id).Delete(userRelation)
	return err
}

//修改
func (this UserRelation) Update(userRelation *UserRelation) error {
	_, err := OrmWeb.AllCols().ID(userRelation.Id).Update(userRelation)
	return err
}

//查询
func (this UserRelation) GetById(id int64) (*UserRelation, error) {
	userRelation := new(UserRelation)
	has, err := OrmWeb.Id(id).Get(userRelation)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userRelation, nil
}

func (this UserRelation) QueryByUserRelation(userRelation *UserRelation) ([]*UserRelation, error) {
	userRelationList := make([]*UserRelation, 0)

	err := OrmWeb.Find(&userRelationList, userRelation)

	if err != nil {
		return nil, err
	}
	return userRelationList, nil
}
