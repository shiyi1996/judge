package models

import (
	. "self/commons/store"
)

type RankPerson struct {
	Id       int64 //个人排行id
	UserId   int64 //用户id
	AcNumber int   //AC数量
	Grade    int   //得分
}

//增加
func (this RankPerson) Add(rankPerson *RankPerson) (int64, error) {
	_, err := OrmWeb.Insert(rankPerson)
	if err != nil {
		return 0, err
	}
	return rankPerson.Id, nil
}

//删除
func (this RankPerson) Remove(id int64) error {
	rankPerson := &RankPerson{}
	_, err := OrmWeb.Id(id).Delete(rankPerson)
	return err
}

//修改
func (this RankPerson) Update(rankPerson *RankPerson) error {
	_, err := OrmWeb.AllCols().ID(rankPerson.Id).Update(rankPerson)
	return err
}

//查询
func (this RankPerson) GetById(id int64) (*RankPerson, error) {
	rankPerson := new(RankPerson)

	has, err := OrmWeb.Id(id).Get(rankPerson)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return rankPerson, nil
}

func (this RankPerson) QueryByRankPerson(rankPerson *RankPerson) ([]*RankPerson, error) {
	rankPersonList := make([]*RankPerson, 0)

	err := OrmWeb.Find(&rankPersonList, rankPerson)

	if err != nil {
		return nil, err
	}
	return rankPersonList, nil
}
