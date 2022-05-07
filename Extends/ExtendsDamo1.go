package Extends

import "fmt"

/*
	项目文件
			作用：编写学生考试系统
*/

type Student struct {
	Name      string
	Classroom int
	Score     float64
}

type LittleStudent struct {
	Student
}

type CollageStudent struct {
	Student
	Credit int
}

// 先是成绩
func (s *Student) ShowInfo() {

	fmt.Printf("name:%v,班级:%v,成绩:%v", s.Name, s.Classroom, s.Score)

}

// 设置成绩
func (s *Student) SetScore(score float64) {

	if score <= 0 {
		return
	}
	s.Score += score

}
