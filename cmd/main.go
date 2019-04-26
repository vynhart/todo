package main

import (
	"github.com/vynhart/todo"
)

func main() {
	server := todo.Server{}
	server.Start(":8080")
}
