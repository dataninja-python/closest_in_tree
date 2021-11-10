package main

import (
	"closest_in_tree/bst"
	// "fmt"
	// "os"
	"testing"
)

func createBST(a []int, t *bst.BST) bst.BST {
	for _, v := range a {
		t.InsertC(v)
	}
}

func TestNodes(t *testing.T) {
	treeSlice1 := []int{10, 5, 15, 2, 5, 13, 22, 1, 14}
	// treeSlice2 := []int{4, 1, 2, 3, 5, 6, 7, 8, 9}

	var tree1 bst.BTree
	tree1 = createBST(treeSlice1, &tree1)

}
