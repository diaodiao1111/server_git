package main

import "awesomeProject/server_git/Extends"

func main() {
	//var res = Extends.Student{
	//	Name: "zhangsan",
	//	Classroom: 10,
	//	Score: 0,
	//}
	//
	//
	//res.SetScore(100)
	//
	//res.ShowInfo()

	var resColl = Extends.CollageStudent{}
	resColl.Name = "zhangsan"
	resColl.ShowInfo()

}
