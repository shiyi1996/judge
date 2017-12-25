package models

import (
	. "self/commons/store"
)

type Group struct {
	Id          int64  //团队表id
	Picture     string //头像
	Name        string //团队名称
	Description string //团队描述
	SetUp       string //设置
}

//增加
func (this Group) Create(group *Group) (int64, error) {
	_, err := OrmWeb.Insert(group)
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

//删除
func (this Group) Remove(id int64) error {
	group := Group{}
	_, err := OrmWeb.Id(id).Delete(group)
	return err
}

//修改
func (this Group) Update(group *Group) error {
	_, err := OrmWeb.Id(group.Id).Update(group)
	return err
}

//查询
func (this Group) GetById(id int64) (*Group, error) {
	group := new(Group)
	has, err := OrmWeb.Id(id).Get(group)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return group, nil
}

func (this Group) GetByName(name string) (*Group, error) {
	group := new(Group)
	has, err := OrmWeb.Where("name = ?", name).Get(group)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return group, nil
}
