package btree

import "github.com/andersfylling/snowflake"

type (
	Node struct {
		Left, Right *Node
		Value       *TreeValue
	}

	BinaryTree struct {
		root *Node
	}

	TreeValue struct {
		key  snowflake.Snowflake // Assume snowflake?
		item interface{}
	}
)

func NewNode(val *TreeValue) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Value: val,
	}
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

func (b *BinaryTree) addIterative(root *Node, val *TreeValue) *Node {
	newNode := NewNode(val)

	if b.root == nil {
		b.root = newNode
	}

	x := root
	var y *Node

	for x != nil {
		y = x
		if val.key < x.Value.key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	if y == nil {
		y = newNode
	} else if val.key < y.Value.key {
		y.Left = newNode
	} else {
		y.Right = newNode
	}

	return y
}

func (b *BinaryTree) Insert(val *TreeValue) {
	//b.root = b.addRecursive(b.root, val)
	b.addIterative(b.root, val)
}

func (b *BinaryTree) findIterative(root *Node, key snowflake.Snowflake) interface{} {
	x := root
	for x != nil {
		if key > x.Value.key {
			x = x.Right
		} else if key < x.Value.key {
			x = x.Left
		} else {
			return x.Value.item
		}
	}
	return nil
}

func (b *BinaryTree) Find(key snowflake.Snowflake) interface{} {
	return b.findIterative(b.root, key)
}

func (b *BinaryTree) deleteIterative(root *Node, key snowflake.Snowflake) *Node {
	curr := root
	var prev *Node

	for curr != nil && curr.Value.key != key {
		prev = curr
		if key < curr.Value.key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	if curr == nil {
		return nil
	}

	if curr.Left == nil || curr.Right == nil {
		var newCurr *Node

		if curr.Left == nil {
			newCurr = curr.Right
		} else {
			newCurr = curr.Left
		}

		if prev == nil {
			b.root = nil
			return nil
		}

		if curr == prev.Left {
			prev.Left = newCurr
		} else {
			prev.Right = newCurr
		}

	} else {
		var p, temp *Node
		temp = curr.Right
		for temp.Left != nil {
			p = temp
			temp = temp.Left
		}

		if p != nil {
			p.Left = temp.Right
		} else {
			curr.Right = temp.Right
		}

		curr.Value = temp.Value
	}

	return curr
}

func (b *BinaryTree) Delete(key snowflake.Snowflake) {
	b.deleteIterative(b.root, key)
}
