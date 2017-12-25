/**
 * Created by shiyi on 2017/11/22.
 * Email: shiyi@fightcoder.com
 */

package models

import (
	"testing"
)

func TestProblemCreate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Titile: "sadas", Description: "1111"}
	if _, err := problem.Create(problem); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestProblemRemove(t *testing.T) {
	InitAllInTest()

	var problem Problem
	if err := problem.Remove(6); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestProblemUpdate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Titile: "sadas", Description: "asdasdasd"}
	if err := problem.Update(problem); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestProblemGetById(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Titile: "sadas", Description: "fffff"}
	Problem{}.Create(problem)

	getProblem, err := Problem{}.GetById(problem.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblem != *problem {
		t.Error("GetById() failed:", "%v != %v", problem, getProblem)
	}
}
func TestProblemQueryByTitile(t *testing.T) {
	InitAllInTest()

	problem1 := &Problem{Titile: "测试"}
	problem2 := &Problem{Titile: "测试"}

	Problem{}.Create(problem1)
	Problem{}.Create(problem2)

	problemList, err := Problem{}.QueryByTitile("测试")
	if err != nil {
		t.Error("QueryByTitile() failed:", err)
	}
	if len(problemList) != 2 {
		t.Error("QueryByTitile() failed:", "count is wrong!")
	}
}
func TestProblemGetByUserId(t *testing.T) {
	InitAllInTest()

	problem := &Problem{UserId: 20}
	problem.Create(problem)

	getProblem, err := problem.GetByUserId(20)
	if err != nil {
		t.Error("GetByUserId() failed:", err)
	}
	if getProblem.UserId != 20 {
		t.Error("GetByUserId() failed:", "%v != %v", problem, getProblem)
	}
}
