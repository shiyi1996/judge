package models

import (
	. "self/commons/store"
)

type RankGroup struct {
	Id       int64   //团队排行id
	GroupId  int64   //团队id
	AcNumber float32 //AC数量
	Grade    float32 //得分
}

//增加
func (this RankGroup) Add(rankGroup *RankGroup) (int64, error) {
	_, err := OrmWeb.Insert(rankGroup)
	if err != nil {
		return 0, err
	}
	return rankGroup.Id, nil
}

//删除
func (this RankGroup) Remove(id int64) error {
	rankGroup := &RankGroup{}
	_, err := OrmWeb.Id(id).Delete(rankGroup)
	return err
}

//修改
func (this RankGroup) Update(rankGroup *RankGroup) error {
	_, err := OrmWeb.AllCols().ID(rankGroup.Id).Update(rankGroup)
	return err
}

//查询
func (this RankGroup) GetById(id int64) (*RankGroup, error) {
	rankGroup := new(RankGroup)

	has, err := OrmWeb.Id(id).Get(rankGroup)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return rankGroup, nil
}

func (this RankGroup) QueryByRankGroup(rankGroup *RankGroup) ([]*RankGroup, error) {
	rankGroupList := make([]*RankGroup, 0)

	err := OrmWeb.Find(&rankGroupList, rankGroup)

	if err != nil {
		return nil, err
	}
	return rankGroupList, nil
}
