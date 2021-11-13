package main

import (
	"closest_in_tree/bst"
	"fmt"
	"log"

	// "sort"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/* func getSlice(a []int, node *bst.BST) {
	for _, v := range a {
		node.Insert(v)
	}
} */

func main() {
	defer timeTrack(time.Now(), "timeTrack")

	// n := bst.NewBST(10)
	n1 := bst.NewBST(100)
	// o := bst.NewBST(4)
	// n := bst.NewBST(10)
	// n.Left = bst.NewBST(5)
	// n.Right = bst.NewBST(15)
	// n.Left.Left = bst.NewBST(2)
	// n.Left.Right = bst.NewBST(5)
	// n.Right.Left = bst.NewBST(13)
	// n.Right.Right = bst.NewBST(22)
	// n.Left.Left.Left = bst.NewBST(1)
	// n.Right.Left.Left = bst.NewBST(14)

	// bTreeSlice := []int{5, 15, 2, 5, 13, 22, 1, 14}
	bTreeSlice1 := []int{502, 55000, 1001, 4500, 204, 205, 207, 208, 206, 203,
		5, 15, 22, 57, 60, 5, 2, 3, 1, 1, 1, 1, 1, -51, -403}
	// bTreeSlice = append(bTreeSlice)
	// sort.Ints(bTreeSlice)
	/* for _, v := range bTreeSlice {
		n.Insert(v)
	} */

	for _, w := range bTreeSlice1 {
		n1.Insert(w)
	}
	//bTreeSliceb := []int{1, 2, 3, 5, 6, 7, 8, 9}
	// sort.Ints(bTreeSliceb)
	/*for _, w := range bTreeSliceb {
		o.Insertb(w)
	}*/
	// bst.PrintTree(os.Stdout, n, 0, 'M')
	// bst.PrintTree(os.Stdout, n1, 0, 'M')
	// fmt.Println("Create tree using normal recusion: ", bst.DFRec(n))
	// bst.DFRec(n)
	// fmt.Printf("\n")
	// bst.DF(n)
	// bst.DFRecImproved(n)
	// bst.DFRec(n)
	// fmt.Println("Create tree using loop and stack: ", bst.DF(n))
	// bst.DF(n)
	// fmt.Printf("\n")
	// fmt.Println("Create tree using loop and stack: ", bst.DF(o))
	// bst.DF(o)
	// fmt.Printf("\n")
	/*fmt.Printf("\n")
	fmt.Println("Mininum number in n tree: ", n.FindMin())
	fmt.Println("Maximum number in n tree: ", n.FindMax())
	bst.PrintTree(os.Stdout, n, 0, 'M')
	fmt.Printf("\n")
	fmt.Println("Mininum number in o tree: ", o.FindMin())
	fmt.Println("Maximum number in o tree: ", o.FindMax())
	bst.PrintTree(os.Stdout, o, 0, 'M')
	fmt.Printf("\n")
	*/
	// p, _ := n.FindNodeWithValue(4)
	// fmt.Println(p.Value)
	// bst.PrintTree(os.Stdout, p, 0, 'M')
	// t := 12
	// fmt.Println("Target to find closest number to in tree below: ", t)
	// fmt.Println(n.FindClosestValue(t))
	// bst.PrintTree(os.Stdout, n, 0, 'M')

	t1 := 29749
	fmt.Println("Target to find closest number to in tree below: ", t1)
	fmt.Println(n1.FindClosestValue(t1))
	// bst.PrintTree(os.Stdout, n1, 0, 'M')
}
