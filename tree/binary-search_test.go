package binTree

import (
	"../util"
	"testing"
)

func TestInsert(t *testing.T) {
	list := util.ObterItens(100, "int")
	newTree := NewTree(list)

	for i, v := range list {
		if i > 0 && !checkBST(newTree.root, 0, 0) {
			t.Errorf("Incorrect binary search tree structure, %+v is incorrectly positioned related to %+v", list[i-1], list[i])
		}

		if ok := newTree.Insert(v); !ok && !util.IsNil(v) {
			t.Errorf("Error inserting non nil item %v", v)
		}
	}
}
