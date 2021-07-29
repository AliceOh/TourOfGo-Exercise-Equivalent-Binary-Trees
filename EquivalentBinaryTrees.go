package main
 
import (
	"fmt"
	"golang.org/x/tour/tree"
)
 
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}
 
func WalkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkRecursive(t.Left, ch)
		ch <- t.Value
		//fmt.Println(t.Value)
		WalkRecursive(t.Right, ch)
	}
}
 
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for n1 := range ch1 {
		n2 := <-ch2
		if n1 != n2 {
			return false
		}
	}
	return true
 
}
 
func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for n := range ch {
		fmt.Println(n)
	}
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
