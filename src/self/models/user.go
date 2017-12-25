package models

import (
	. "self/commons/store"
)

type User struct {
	Id           int64
	AccountId    int64  //账号Id
	NickName     string //昵称
	Description  string //个人描述
	Sex          int    //性别
	Birthday     int64  //生日
	DailyAddress string //日常所在地：省、市
	RecvAddress  string //收件地址，仅自己可见
	TShirtSize   string //T-恤尺码(S、M、L、XL、XXL、XXL)
	StatSchool   int    //当前就学状态(小学及以下、中学学生、大学学生、非在校生)
	Blog         string //博客地址
	Git          string //Git地址
	Avator       string //头像
}

//增加
func (this User) Create(user *User) (int64, error) {
	_, err := OrmWeb.Insert(user) //第一个参数为影响的行数
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

//删除
func (this User) Remove(id int64) error {
	user := User{}
	_, err := OrmWeb.Id(id).Delete(user)
	return err
}

//修改
func (this User) Update(user *User) error {
	_, err := OrmWeb.AllCols().ID(user.Id).Update(user)
	return err
}

//查询
func (this User) GetById(id int64) (*User, error) {
	user := new(User)
	has, err := OrmWeb.Id(id).Get(user) //第一个为 bool 类型，表示是否查找到记录

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func (this User) GetByAccountId(accountId int64) (*User, error) {
	user := new(User)
	has, err := OrmWeb.Where("account_id = ?", accountId).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func (this User) QueryByName(nickname string) ([]*User, error) {
	userList := make([]*User, 0)
	err := OrmWeb.Where("nick_name like ?", "%"+nickname+"%").Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
