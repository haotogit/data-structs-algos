package binTree

import (
	"../util"
	"testing"
)

func TestInsert(t *testing.T) {
	var list []interface{}
	newTree := NewTree(list)
	list = util.ObterItens(100, "int")

	for i, v := range list {
		if i > 0 && !checkBST(newTree.root, 0, 0) {
			t.Errorf("Incorrect binary search tree structure, %+v is incorrectly positioned related to %+v", list[i-1], list[i])
		}

		if ok := newTree.Insert(v); !ok && !util.IsNil(v) {
			t.Errorf("Error inserting non nil item %v", v)
		}
	}
}

func TestSize(t *testing.T) {
	for i := 0; i < 50; i++ {
		list := util.ObterItens(50, "int")
		totalSize := len(list)
		newTree := NewTree(list)
		if totalSize != newTree.root.size {
			t.Errorf("Incorrect tree size expected %d, but got %d", totalSize, newTree.root.size)
		}
	}
}

func TestSearch(t *testing.T) {
	list := util.ObterItens(100, "int")
	newTree := NewTree(list)

	for i := 0; i < len(list); i += 5 {
		curr := list[i]
		found := newTree.Search(newTree.root, curr).val.(int)

		if curr != found {
			t.Errorf("Found wrong item expected %d but got %d", curr, found)
		}
	}
}

func TestDelete(t *testing.T) {
	list := util.ObterItens(100, "int")
	newTree := NewTree(list)

	for i := 0; i < len(list); i += 10 {
		prevSize := newTree.root.size
		newTree.Delete(list[i])
		found := newTree.Search(newTree.root, list[i])

		if prevSize == newTree.root.size {
			t.Errorf("Incorrect size after deletion, expected %d, but got %d", prevSize-1, newTree.root.size)
		}

		if found != nil {
			t.Errorf("Did not delete %d from tree", list[i])
		}

		if !checkBST(newTree.root, 0, 0) {
			t.Errorf("Invalid bst structure after deletion of %d", list[i])
		}
	}
}
