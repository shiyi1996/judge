package models

import (
	. "self/commons/store"
)

type Contest struct {
	Id          int64  //比赛id
	UserId      int64  //举办者id
	Type        int    //比赛类型
	Name        string //比赛名称
	Description string //比赛描述
	ProblemList string //题目列表
	StartTime   int64  //开始时间
	FrozenTime  int64  //封榜时间
	EndTime     int64  //结束时间
	Password    string //密码，为空表示公开
}

//增加
func (this Contest) Create(contest *Contest) (int64, error) {
	_, err := OrmWeb.Insert(contest)
	if err != nil {
		return 0, err
	}
	return contest.Id, nil
}

//删除
func (this Contest) Remove(id int64) error {
	contest := Contest{}
	_, err := OrmWeb.Id(id).Delete(contest)
	return err
}

//修改
func (this Contest) Update(contest *Contest) error {
	_, err := OrmWeb.Id(contest.Id).Update(contest)
	return err
}

//查询
func (this Contest) GetById(id int64) (*Contest, error) {
	contest := new(Contest)
	has, err := OrmWeb.Id(id).Get(contest)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return contest, nil
}

func (this Contest) GetByUserId(userId int64) (*Contest, error) {
	contest := new(Contest)
	has, err := OrmWeb.Where("user_id = ?", userId).Get(contest)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return contest, nil
}

func (this Contest) QueryByName(name string) ([]*Contest, error) {
	userList := make([]*Contest, 0)
	err := OrmWeb.Where("name like ?", "%"+name+"%").Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (this Contest) QueryByType(t int64) ([]*Contest, error) {
	contestList := make([]*Contest, 0)
	err := OrmWeb.Where("type = ?", t).Find(&contestList)
	if err != nil {
		return nil, err
	}
	return contestList, nil
}
