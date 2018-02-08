package avl

import (
	"testing"
)

var keys1 = []int{10, 20, 25}
var keys2 = []int{10, 20, 25, 30}
var keys3 = []int{10, 30, 25, 20}

func TestTree_Insert(t *testing.T) {
	tree := New()
	for _, v := range keys1 {
		tree.Insert(v)
	}
	root := tree.GetRoot()
	rootRight := root.GetRight()
	rootLeft := root.GetLeft()

	if root.GetParent() != nil {
		t.Fatal("Root parent must be nil")
	}

	if root.GetKey() != 10 {
		t.Fatalf("Root key expect %d instead %d", 10, tree.GetRoot().GetKey())
	}

	if rootRight.GetKey() != 20 {
		t.Fatalf("Root right expect %d instead %d", 20, rootRight.GetKey())
	}

	if rootLeft != nil {
		t.Fatal("Root left must be nil")
	}
}

func TestTree_Insert2(t *testing.T) {
	tree := New()
	for _, v := range keys2 {
		tree.Insert(v)
	}
	root := tree.GetRoot()
	rootRight := root.GetRight()
	rootLeft := root.GetLeft()

	if root.GetParent() != nil {
		t.Fatal("Root parent must be nil")
	}

	if root.GetKey() != 20 {
		t.Fatalf("Root key expect %d instead %d", 20, tree.GetRoot().GetKey())
	}

	if rootRight.GetKey() != 25 {
		t.Fatalf("Root right expect %d instead %d", 25, rootRight.GetKey())
	}

	if rootLeft.GetKey() != 10 {
		t.Fatalf("Root left expect %d instead %d", 10, rootLeft.GetKey())
	}

	if rootLeft.GetRight()!= nil {
		t.Fatal("Root left right must be nil")
	}

	if rootLeft.GetParent() != root {
		t.Fatal("Root left parent must be root")
	}
}

func TestTree_Insert3(t *testing.T) {
	tree := New()
	for _, v := range keys3 {
		tree.Insert(v)
	}
	root := tree.GetRoot()
	rootRight := root.GetRight()
	rootLeft := root.GetLeft()

	if root.GetParent() != nil {
		t.Fatal("Root parent must be nil")
	}

	if root.GetKey() != 25 {
		t.Fatalf("Root key expect %d instead %d", 25, tree.GetRoot().GetKey())
	}

	if rootRight.GetKey() != 30 {
		t.Fatalf("Root right key expect %d instead %d", 30, tree.GetRoot().GetKey())
	}

	if rootLeft.GetKey() != 10 {
		t.Fatalf("Root left expect %d instead %d", 10, rootLeft.GetKey())
	}

	if rootLeft.GetRight().GetKey() != 20 {
		t.Fatalf("Root left right expect %d instead %d", 20, rootLeft.GetRight().GetKey())
	}
}