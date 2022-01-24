package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkRecursive(t.Left, ch)
	ch <- t.Value
	WalkRecursive(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Inorder traversal
	WalkRecursive(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		j := <-ch2
		if i != j {
			return false
		}
	}

	return true
}

func main() {
	// Test walk
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(2)))
}
