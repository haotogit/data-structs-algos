package binTree

import (
	"github.com/haotogit/data-structs-algos/queue"
	"github.com/haotogit/data-structs-algos/util"
	"fmt"
)

type btNode struct {
	key   uint32
	val   interface{}
	left  *btNode
	right *btNode
	size  int
}

type binTreee interface {
	Insert(x interface{}) bool
	Height(currRoot *btNode) int
	Search(x interface{}) interface{}
	Delete(x interface{}) bool
	Size(currRoot *btNode) int
	add(currRoot, currNode *btNode) btNode
	addAll(list []interface{})
	remove(currRoot *btNode, x interface{}) *btNode
	min(currRoot btNode) btNode
	inOrder(currRoot *btNode)
	levelOrder()
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

	currRoot.size = 1 + bt.Size(currRoot.left) + bt.Size(currRoot.right)
	return currRoot
}

func (bt binTree) Height(currNode *btNode) int {
	if currNode == nil {
		return -1
	}

	return 1 + util.GetGreatest(bt.Height(currNode.left), bt.Height(currNode.right))
}

func (bt binTree) Search(curr *btNode, x interface{}) *btNode {
	if curr == nil || util.IsNil(x) {
		return nil
	}

	compared := util.Greater(uint32(x.(int)), curr.key)
	if compared == 0 {
		return curr
	} else if compared == 1 {
		return bt.Search(curr.right, x)
	} else {
		return bt.Search(curr.left, x)
	}
}

func (bt *binTree) Delete(x interface{}) bool {
	if util.IsNil(x) || bt.root == nil {
		return false
	}

	bt.root = bt.remove(bt.root, x)
	return true
}

func (bt binTree) Size(currNode *btNode) int {
	if currNode == nil {
		return 0
	}

	return currNode.size
}

func (bt *binTree) remove(currRoot *btNode, x interface{}) *btNode {
	if currRoot == nil {
		return nil
	}

	compared := util.Greater(currRoot.key, uint32(x.(int)))
	if compared == 1 {
		currRoot.left = bt.remove(currRoot.left, x)
	} else if compared == -1 {
		currRoot.right = bt.remove(currRoot.right, x)
	} else {
		// if one child
		if currRoot.right == nil {
			return currRoot.left
		} else if currRoot.left == nil {
			return currRoot.right
		}

		min := bt.min(currRoot.right)
		// replace node to be removed with min node
		currRoot.key = min.key
		currRoot.val = min.val
		// and set min right child to it's previous spot
		currRoot.right = bt.remove(currRoot.right, int(currRoot.key))
	}

	currRoot.size = 1 + bt.Size(currRoot.left) + bt.Size(currRoot.right)
	return currRoot
}

func (bt binTree) min(root *btNode) *btNode {
	if root.left == nil {
		return root
	}

	return bt.min(root.left)
}

func (bt *binTree) addAll(list []interface{}) {
	for _, item := range list {
		bt.Insert(item)
	}
}

func (bt binTree) inOrder(currRoot *btNode) {
	if currRoot != nil {
		bt.inOrder(currRoot.left)
		fmt.Printf(">> %+v\n", currRoot)
		bt.inOrder(currRoot.right)
	}
}

func (bt binTree) levelOrder() {
	newQ := queue.NewQ(bt.root.size)
	newQ.Enqueue(bt.root)

	for newQ.Size() != 0 {
		curr := (newQ.Dequeue()).(*btNode)
		fmt.Printf(">> %+v\n", curr)
		if curr.left != nil {
			newQ.Enqueue(curr.left)
		}

		if curr.right != nil {
			newQ.Enqueue(curr.right)
		}
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

func (bt binTree) newNode(item interface{}) *btNode {
	var currKey uint32
	// TODO fix this
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
