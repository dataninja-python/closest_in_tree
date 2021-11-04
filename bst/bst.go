package bst

import (
	"fmt"
	"sync"
	"time"
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

func NewBST(data int) *BST {
	return &BST{data, nil, nil}
}

func (n *BST) Process() {
	time.Sleep(2 * time.Second) // add time consuming things
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
	defer wg.Done()

	if n == nil {
		return
	}
	wg.Add(1)
	go DFRecImproved(n.Left)
	n.Process()
	wg.Add(1)
	go DFRecImproved(n.Right)
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
			current, _ = s.Pop()
			current.Process()
			current = current.Right
		}
	}
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return -1
}

// NOTE: I have never really been exposed to binary trees or graphs except as
// objects or databases; so I learned initially from these implementations:
// https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
// https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php
// https://medium.com/analytics-vidhya/how-to-speed-up-a-tree-structure-traversal-in-go-cd4bd6775520