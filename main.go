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
	start := time.Now()

	n := NewNode(10)
	n.Left = NewNode(5)

	Insert(5).
		Insert(15).
		Insert(2).
		Insert(5).
		Insert(13).
		Insert(22).
		Insert(1).
		Insert(14)
	// bst.PrintTree(os.Stdout, aTree.Top, 0, 'M')
	fmt.Println(aTree.Top.FindClosestValue(12))
}
