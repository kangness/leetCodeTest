package main

import (
	"fmt"
	"time"
	"math/rand"
	. "./tree"
)

func main() {
	var node, isFind  *BaseTree
	manager := AVLTree{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 50 ; i ++ {
		num := rand.Intn(100)
		isFind = manager.Find(node, num)
		if isFind != nil {
			//fmt.Printf("%+v", isFind)
			continue
		}
		node = manager.InsertNode(node, num)
	}
	fmt.Println(manager.LDR(node))
}
