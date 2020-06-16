package binTree

import (
	"../util"
)

type btNode struct {
	key   uint32
	val   interface{}
	left  *btNode
	right *btNode
	size  int
}

type binTreee interface {
	Insert(currNode *btNode) bool
	add(currRoot, currNode *btNode) btNode
	addAll(list []interface{})
	checkBST(currNode *btNode, min, max uint32) bool
}

type binTree struct {
	root *btNode
}

func (bt *binTree) Insert(item interface{}) bool {
	if util.IsNil(item) {
		return false
	}

	bt.root = bt.add(bt.root, bt.newNode(item))
	return true
}

func (bt *binTree) add(currRoot, freshNode *btNode) *btNode {
	if currRoot == nil {
		return freshNode
	}

	isLesser := util.Greater(currRoot.key, freshNode.key)
	if isLesser == 1 {
		if currRoot.left != nil {
			currRoot.left = bt.add(currRoot.left, freshNode)
		} else {
			currRoot.left = freshNode
		}
	} else if isLesser == -1 {
		if currRoot.right != nil {
			currRoot.right = bt.add(currRoot.right, freshNode)
		} else {
			currRoot.right = freshNode
		}
	}

	currRoot.size = 1 + size(currRoot.left) + size(currRoot.right)
	return currRoot
}

func (bt *binTree) addAll(list []interface{}) {
	for _, item := range list {
		bt.Insert(item)
	}
}

func checkBST(currNode *btNode, min, max uint32) bool {
	if currNode == nil {
		return true
	}

	if (!util.IsNil(min) &&
		util.Greater(currNode.key, min) <= 0) ||
		(!util.IsNil(max) &&
			util.Greater(currNode.key, max) >= 0) {
		return false
	}

	// ensures currNode.key is less|greater than parent
	return checkBST(currNode.left, min, currNode.key) &&
		checkBST(currNode.right, currNode.key, max)
}

func size(currNode *btNode) int {
	if currNode == nil {
		return 0
	}

	return currNode.size
}

func (bt binTree) newNode(item interface{}) *btNode {
	var currKey uint32
	// todo fix this
	if util.GetType(item) == "string" {
		currKey = uint32(bt.root.size + 1)
	} else {
		currKey = uint32(item.(int))
	}

	xNode := new(btNode)
	xNode.key = currKey
	xNode.val = item
	xNode.size = 1
	return xNode
}

func NewTree(list []interface{}) *binTree {
	newTree := &binTree{}
	newTree.addAll(list)
	return newTree
}
