package main

import (
	"closest_in_tree/bst"
	"fmt"
	"os"
	"testing"
)

func TestNodes(t *testing.T) {
	var gotTree bst.Root
	gotTree.Insert(10).
		Insert(5).
		Insert(15).
		Insert(2).
		Insert(5).
		Insert(13).
		Insert(22).
		Insert(1).
		Insert(14)

	bst.PrintTree(os.Stdout, gotTree.Top, 0, 'M')
	fmt.Println(gotTree.Top.FindClosestValue(12))
}
