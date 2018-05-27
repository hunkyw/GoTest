package main

import "fmt"

type treenode struct {
	value int
	left,right *treenode
} 

func main() {
	var  root  treenode

	root = treenode{value:3}
	root.left = &treenode{}
	root.right = &treenode{5,nil,nil}
	root.right.left = new(treenode)

	nodes := []treenode{
		{value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
}
