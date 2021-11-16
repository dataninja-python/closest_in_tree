package bst

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var wg sync.WaitGroup

type Stack []BST

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(n *BST) {
	*s = append(*s, *n)
}

func (s *Stack) POP() (*BST, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		top := len(*s) - 1
		item := (*s)[top]
		*s = (*s)[:top]
		return &item, true
	}
}

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

type VAL struct {
	keys  []int
	diffs []int
}

func NewBST(data int) *BST {
	tmp := &BST{Value: data, Left: nil, Right: nil}
	return tmp
}

func (n *BST) Process() {
	fmt.Print(n.Value, " ")
}

// recursively walk through tree
func DFRec(n *BST) {
	if n == nil {
		return
	}
	DFRec(n.Left)
	n.Process()
	DFRec(n.Right)
}

// recursively walk through tree
func DFRecImproved(n *BST) {
	if n == nil {
		return
	}
	wg.Add(1)
	go DFRecImproved(n.Left)
	n.Process()
	wg.Add(1)
	go DFRecImproved(n.Right)
	wg.Done()
}

func DF(n *BST) {
	s := Stack{}
	current := n
	for {
		if s.IsEmpty() && current == nil {
			return
		}
		if current != nil {
			s.Push(current)
			current = current.Left
		} else {
			current, _ = s.POP()
			current.Process()
			current = current.Right
		}
	}
}

func DFOrdered(n *BST) {
	s := Stack{}
	current := n
	for {
		if s.IsEmpty() && current == nil {
			return
		}
		if current != nil {
			s.Push(current)
			current = current.Left
		} else {
			current, _ = s.POP()
			current.Process()
			current = current.Right
		}
	}
}

func (n *BST) Insert(v int) {
	if n == nil {
		return
	} else if v < n.Value {
		if n.Left == nil {
			n.Left = &BST{Value: v, Left: nil, Right: nil}
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &BST{Value: v, Left: nil, Right: nil}
		} else {
			n.Right.Insert(v)
		}
	}
}

func PrintTree(w io.Writer, n *BST, ns int, ch rune) {
	if n == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, n.Value)
	PrintTree(w, n.Left, ns+2, 'L')
	PrintTree(w, n.Right, ns+2, 'R')
}

func (n *BST) InsertSlice(aSlice []int) {
	for _, v := range aSlice {
		n.Insert(v)
	}
}

func (n *BST) Insertb(value int) error {
	if n == nil {
		return errors.New("Tree is nil")
	}

	if n.Value > value {
		if n.Left == nil {
			n.Left = &BST{Value: value}
			return nil
		}
		return n.Left.Insertb(value)
	} else {
		if n.Right == nil {
			n.Right = &BST{Value: value}
			return nil
		}
		return n.Right.Insertb(value)
	}
}

func (n *BST) FindMin() int {
	if n.Left == nil {
		return n.Value
	}
	return n.Left.FindMin()
}

func (n *BST) FindMax() int {
	if n.Right == nil {
		return n.Value
	}
	return n.Right.FindMax()
}

func (n *BST) PrintInOrder() {
	if n == nil {
		return
	}
	n.Left.PrintInOrder()
	fmt.Print(n.Value)
	n.Right.PrintInOrder()
}

// returns node that contains the supplied int param
func (n *BST) FindNodeWithValue(v int) (BST, bool) {
	// if tree is empty return empty tree and false
	if n == nil {
		return BST{}, false
	}

	// select from multiple options each time through recursion
	switch {
	// if the current node contains the value return it and true
	// else keep recursively going through tree
	case v == n.Value:
		return *n, true
	case v < n.Value:
		return n.Left.FindNodeWithValue(v)
	default:
		return n.Right.FindNodeWithValue(v)
	}
}

// returns the int and bool to say if first int param < second one
// can use recursively in binary search to check walk through tree
// Note: core check is if x < y, so if y <= x will return y
func (n *BST) IsXLessThanY(x, y int) bool {
	if x < y {
		return true
	}
	// this capture if the values are equal
	// child >= parent = right is how the tree is constructed
	// thus, this is the proper way to execut this
	return false
}

func getAllNodeValues(aTree *BST, theTarget int, theMap map[int]bool) map[int]bool {
	switch {
	case aTree.Value == theTarget:
		if _, ok := theMap[aTree.Value]; !ok {
			theMap[aTree.Value] = true
		}
		return theMap
	case theTarget < aTree.Value:
		if aTree.Left == nil {
			if _, ok := theMap[aTree.Value]; !ok {
				theMap[aTree.Value] = true
			}
			return theMap
		}
		getAllNodeValues(aTree, theTarget, theMap)
	default:
		if aTree.Right == nil {
			if _, ok := theMap[aTree.Value]; !ok {
				theMap[aTree.Value] = true
			}
			return theMap
		}
		getAllNodeValues(aTree, theTarget, theMap)
	}
	return theMap
}

func (tree *BST) FindClosestValue(target int) int {
	// pseudocode:
	// walk every node that the new value would walk to fit in the tree
	emptyMap := make(map[int]bool)
	nodeMap := getAllNodeValues(tree, target, emptyMap)
	fmt.Println(nodeMap)
	return -1
}

// NOTE: I have not used binary trees or graphs in decades except as
// objects or databases; so I re-learned initially from these implementations:
// https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
// https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php
// https://medium.com/analytics-vidhya/how-to-speed-up-a-tree-structure-traversal-in-go-cd4bd6775520
// levelup.gitconnected.com/binary-seaarch-trees-in-go-58f9126eb36b
