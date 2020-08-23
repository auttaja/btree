package btree

import (
	"github.com/andersfylling/snowflake"
)

type (
	AVLNode struct {
		Left, Right *AVLNode
		Value       *TreeValue
		Height      int
	}

	AVLTree struct {
		root *AVLNode
	}
)

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func NewAVLNode(val *TreeValue) *AVLNode {
	return &AVLNode{
		Value:  val,
		Height: 1,
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func height(n *AVLNode) int {
	if n == nil {
		return 0
	}

	return n.Height
}

func rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	z := x.Right

	x.Right = y
	y.Left = z

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

func leftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	z := y.Left

	y.Left = x
	x.Right = z

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

func getBalance(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func (t *AVLTree) recursiveInsert(node *AVLNode, value *TreeValue) *AVLNode {
	if node == nil {
		return NewAVLNode(value)
	}

	if value.key < node.Value.key {
		node.Left = t.recursiveInsert(node.Left, value)
	} else if value.key > node.Value.key {
		node.Right = t.recursiveInsert(node.Right, value)
	} else {
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))

	balance := getBalance(node)

	if balance > 1 && value.key < node.Left.Value.key {
		return rightRotate(node)
	}

	if balance < -1 && value.key > node.Right.Value.key {
		return leftRotate(node)
	}

	if balance > 1 && value.key > node.Left.Value.key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && value.key < node.Right.Value.key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func (t *AVLTree) Insert(value *TreeValue) {
	t.root = t.recursiveInsert(t.root, value)
}

func minValueNode(node *AVLNode) *AVLNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}

	return current
}

func (t *AVLTree) recursiveDelete(root *AVLNode, key snowflake.Snowflake) *AVLNode {
	if root == nil {
		return root
	}

	if key < root.Value.key {
		root.Left = t.recursiveDelete(root.Left, key)
	} else if key > root.Value.key {
		root.Right = t.recursiveDelete(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			var temp *AVLNode
			if root.Left != nil {
				temp = root.Left
			} else if root.Right != nil {
				temp = root.Right
			}

			if temp == nil {
				temp = root
				root = nil
			} else {
				root = temp
			}
		} else {
			temp := minValueNode(root.Right)
			root.Value = temp.Value
			root.Right = t.recursiveDelete(root.Right, temp.Value.key)
		}
	}

	if root == nil {
		return root
	}

	root.Height = 1 + max(height(root.Left), height(root.Right))

	balance := getBalance(root)

	if balance > 1 && getBalance(root.Left) >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && getBalance(root.Left) < 0 {
		root.Left = leftRotate(root.Left)
		return rightRotate(root)
	}

	if balance < -1 && getBalance(root.Right) <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && getBalance(root.Right) > 0 {
		root.Right = rightRotate(root.Right)
		return leftRotate(root)
	}

	return root
}

func (t *AVLTree) Delete(key snowflake.Snowflake) {
	t.root = t.recursiveDelete(t.root, key)
}

func (t *AVLTree) recursiveFind(root *AVLNode, key snowflake.Snowflake) *AVLNode {
	if root == nil || root.Value.key == key {
		return root
	}

	if root.Value.key < key {
		return t.recursiveFind(root.Right, key)
	}

	return t.recursiveFind(root.Left, key)
}

func (t *AVLTree) Find(key snowflake.Snowflake) interface{} {
	node := t.recursiveFind(t.root, key)
	if node == nil {
		return nil
	}
	return node.Value.item
}
