package models

import (
	. "self/commons/store"
)

type Member struct {
	Id      int64  //成员表id
	GroupId int64  //团队Id
	UserId  int64  //用户id
	Role    string //角色
}

//增加
func (this Member) Create(member *Member) (int64, error) {
	_, err := OrmWeb.Insert(member)
	if err != nil {
		return 0, err
	}
	return member.Id, nil
}

//删除
func (this Member) Remove(id int64) error {
	member := Member{}
	_, err := OrmWeb.Id(id).Delete(member)
	return err
}

//修改
func (this Member) Update(member *Member) error {
	_, err := OrmWeb.Id(member.Id).Update(member)
	return err
}

//查询
func (this Member) GetById(id int64) (*Member, error) {
	member := new(Member)
	has, err := OrmWeb.Id(id).Get(member)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return member, nil
}

func (this Member) GetByGroupId(groupId int64) (*Member, error) {
	member := new(Member)
	has, err := OrmWeb.Where("group_id = ?", groupId).Get(member)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return member, nil
}
func (this Member) GetByUserId(userId int64) (*Member, error) {
	member := new(Member)
	has, err := OrmWeb.Where("user_id = ?", userId).Get(member)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return member, nil
}

func (this Member) QueryByRole(role string) ([]Member, error) {
	memberList := make([]Member, 0)
	err := OrmWeb.Where("role like ?", "%"+role+"%").Find(&memberList)
	if err != nil {
		return nil, err
	}
	return memberList, nil
}
