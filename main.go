package main

import (
	"closest_in_tree/bst"
<<<<<<< HEAD
=======
	"fmt"

	// "fmt"
>>>>>>> 65ff58360cfd522680705cab10cf2bb910ea8cf7
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
	o := bst.NewBST(4)
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

	bTreeSliceb := []int{1, 2, 3, 5, 6, 7, 8, 9}
	for _, w := range bTreeSliceb {
		o.Insertb(w)
	}
	// bst.PrintTree(os.Stdout, n, 0, 'M')
	bst.DFRec(n)
	// bst.DF(n)
	// bst.DFRecImproved(n)
	// bst.DFRec(n)
	bst.DF(n)
	bst.DF(o)
	fmt.Printf("\n")
	fmt.Println("Mininum number in the tree: ", o.FindMin())
	fmt.Println("Maximum number in the tree: ", o.FindMax())
}
