package models

import (
	. "self/commons/store"
)

type Submit struct {
	Id            int64
	ProblemId     int64  //题目ID
	UserId        int64  //提交用户ID
	Language      string //提交语言
	SubmitTime    int64  //提交时间
	RunningTime   int64  //耗时(ms)
	RunningMemory int64  //所占空间
	Result        string //运行结果
	ResultDes     string //结果描述
	CaseResult    string //测试数据结果
	Code          string //提交代码
}

//增加
func (this Submit) Create(submit *Submit) (int64, error) {
	_, err := OrmWeb.Insert(submit)
	if err != nil {
		return 0, err
	}
	return submit.Id, nil
}

//删除
func (this Submit) Remove(id int64) error {
	submit := Submit{}
	_, err := OrmWeb.Id(id).Delete(submit)
	return err
}

//修改
func (this Submit) Update(submit *Submit) error {
	_, err := OrmWeb.AllCols().ID(submit.Id).Update(submit)
	return err
}

//查询
func (this Submit) GetById(id int64) (*Submit, error) {
	submit := new(Submit)
	has, err := OrmWeb.Id(id).Get(submit)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submit, nil
}

func (this Submit) QueryBySubmit(problemId, userId int64, language, resultDes string) ([]*Submit, error) {
	submit := Submit{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes}
	submitList := make([]*Submit, 0)

	err := OrmWeb.Find(&submitList, submit)

	if err != nil {
		return nil, err
	}
	return submitList, nil
}
