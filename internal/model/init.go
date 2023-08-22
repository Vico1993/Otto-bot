package model

import "fmt"

var ListOfChats map[int64]Chat

func Init() {
	ListOfChats = make(map[int64]Chat)

	fmt.Println("Model Initiated")
}
