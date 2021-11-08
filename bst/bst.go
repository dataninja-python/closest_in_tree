package bst

import (
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

func NewBST(data int) *BST {
	tmp := &BST{Value: data, Left: nil, Right: nil}
	// tmp.AllValues = append(tmp.AllValues, data)
	return tmp
}

func (n *BST) Process() {
	// time.Sleep(2 * time.Second) // add time consuming things
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

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return -1
}

// NOTE: I have never really been exposed to binary trees or graphs except as
// objects or databases; so I learned initially from these implementations:
// https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
// https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php
// https://medium.com/analytics-vidhya/how-to-speed-up-a-tree-structure-traversal-in-go-cd4bd6775520
