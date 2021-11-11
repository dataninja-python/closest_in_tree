package main

import (
	"closest_in_tree/bst"
	"fmt"
	"os"
	"testing"
)

type treeTest struct {
	bstSlice BST
	expected int
}

func createBST(a []int) BST {
	var b bst.BST
	for i, v := range a {
		if i == 0 {
			b = bst.NewBST(v)
		}
		b.Insert(v)
	}
}

func TestNodes(t *testing.T) {
	treeSlice1 := []int{10, 5, 15, 2, 5, 13, 22, 1, 14}
	treeSlice2 := []int{4, 1, 2, 3, 5, 6, 7, 8, 9}

	bst.PrintTree(os.Stdout, gotTree.Top, 0, 'M')
	fmt.Println(gotTree.Top.FindClosestValue(12))
	fmt.Println(gotTree.Top.FindClosestValue(10))
	fmt.Println(gotTree.Top.FindClosestValue(5))
}
