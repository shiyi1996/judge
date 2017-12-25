package models

import (
	. "self/commons/store"
)

type Player struct {
	Id        int64 //参与比赛者表id
	ContestId int64 //竞赛id
	UserId    int64 //用户id
}

//增加
func (this Player) Create(player *Player) (int64, error) {
	_, err := OrmWeb.Insert(player)
	if err != nil {
		return 0, err
	}
	return player.Id, nil
}

//删除
func (this Player) Remove(id int64) error {
	player := Player{}
	_, err := OrmWeb.Id(id).Delete(player)
	return err
}

//修改
func (this Player) Update(player *Player) error {
	_, err := OrmWeb.Id(player.Id).Update(player)
	return err
}

//查询
func (this Player) GetById(id int64) (*Player, error) {
	player := new(Player)
	has, err := OrmWeb.Id(id).Get(player)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return player, nil
}

func (this Player) GetByContestId(contestId int64) (*Player, error) {
	player := new(Player)
	has, err := OrmWeb.Where("contest_id = ?", contestId).Get(player)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return player, nil
}
func (this Player) GetByUserId(userId int64) (*Player, error) {
	player := new(Player)
	has, err := OrmWeb.Where("user_id = ?", userId).Get(player)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return player, nil
}
