package server

var ADDTASK = 0
var LISTTASKS = 1

type Command struct {
	Name      int
	Optionals struct{}
}

type Response struct {
	Content string
}
