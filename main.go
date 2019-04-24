package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkRecursive(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		el1, ok1 := <-ch1
		el2, ok2 := <-ch2
		if !(ok1 && ok2) {
			break
		}
		if el1 != el2 {
			return false
		}
	}
	return true
}

func main() {
	//testWalk()
	testSame()
}

func testSame() {
	if Same(tree.New(1), tree.New(2)) {
		println("Same")
	} else {
		print("Not the same")
	}
}

func testWalk() {
	done := make(chan bool, 1)
	ch := make(chan int, 10)
	tr := tree.New(1)
	println(tr.String())
	go Walk(tr, ch)
	go pr(ch, done)
	<-done
}

func pr(ch chan int, done chan bool) {
	for e := range ch {
		print(e)
	}
	close(done)
}
