package main

import (
	"closest_in_tree/bst"
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "timeTrack")

	n := bst.NewBST(10)
	// n := bst.NewBST(10)
	// n.Left = bst.NewBST(5)
	// n.Right = bst.NewBST(15)
	// n.Left.Left = bst.NewBST(2)
	// n.Left.Right = bst.NewBST(5)
	// n.Right.Left = bst.NewBST(13)
	// n.Right.Right = bst.NewBST(22)
	// n.Left.Left.Left = bst.NewBST(1)
	// n.Right.Left.Left = bst.NewBST(14)

	bTreeSlice := []int{5, 15, 2, 5, 13, 22, 1, 14}
	// bTreeSlice = append(bTreeSlice)
	for _, v := range bTreeSlice {
		n.Insert(v)
	}

	// bst.PrintTree(os.Stdout, n, 0, 'M')
	// bst.DFRec(n)
	bst.DF(n)
	fmt.Println(n.AllValues)
}
