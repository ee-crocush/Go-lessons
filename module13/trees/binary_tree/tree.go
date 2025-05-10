package main

import (
	"fmt"
)

// Узел дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Дерево
type Tree struct {
	root *Node
}

// Добавление элемента в дерево
func (t *Tree) Append(value int) {
	newNode := &Node{data: value}

	if t.root == nil {
		t.root = newNode
		return
	}

	t.root.appendNode(newNode)
}

func (n *Node) appendNode(newNode *Node) {
	if newNode.data < n.data {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.appendNode(newNode)
		}
	} else {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.appendNode(newNode)
		}
	}
}

// Удаление узла
func (t *Tree) Delete(value int) error {
	var deleted bool

	t.root, deleted = t.root.deleteNode(value)

	if !deleted {
		return fmt.Errorf("Узел %d не найден", value)
	}
	return nil
}

func (n *Node) deleteNode(value int) (*Node, bool) {
	deleted := false

	if n == nil {
		return nil, deleted
	}

	if value < n.data { // Удаляем в левой части дерева
		n.left, deleted = n.left.deleteNode(value)
	} else if value > n.data { // Удаляем в правой части дерева
		n.right, deleted = n.right.deleteNode(value)
	} else { // Узел был найден
		if n.left == nil && n.right == nil { // Удаление листа
			return nil, true
		} else if n.left == nil { //Узел с правым потомком
			return n.right, true
		} else if n.right == nil { //Узел с левым потомком
			return n.left, true
		} else { //Узел с двумя потомками
			minRight := n.right.findMin()
			n.data = minRight.data
			n.right, deleted = n.right.deleteNode(minRight.data)
			return n, deleted
		}
	}

	return n, deleted
}

func (n *Node) findMin() *Node {
	if n.left == nil {
		return n
	}
	return n.left.findMin()
}

// Функция поиска узла
func (t *Tree) Search(value int) (*Node, error) {
	return t.root.searchNode(value)
}

func (n *Node) searchNode(value int) (*Node, error) {
	if n == nil {
		return nil, fmt.Errorf("Узел %d не найден", value)
	}
	if n.data == value {
		return n, nil
	} else if value < n.data {
		return n.left.searchNode(value)
	} else {
		return n.right.searchNode(value)
	}
}

// Вывод дерева в ширину
func (t *Tree) ShowWide() {
	if t.root == nil {
		return
	}

	queue := []*Node{t.root}
	for len(queue) > 0 {
		nextQueue := []*Node{}
		for _, node := range queue {
			fmt.Printf("%d ", node.data)
			if node.left != nil {
				nextQueue = append(nextQueue, node.left)
			}
			if node.right != nil {
				nextQueue = append(nextQueue, node.right)
			}
		}
		fmt.Println()
		queue = nextQueue
	}
}

func main() {
	values := []int{10, 5, 7, 16, 13, 2, 20}
	tree := &Tree{}
	for _, v := range values {
		tree.Append(v)
	}

	tree.ShowWide()
	// if utils := tree.Delete(50); utils != nil {
	// 	fmt.Println(utils)
	// } else {
	// 	tree.ShowWide()
	// }

	node, err := tree.Search(130)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Узел найден: %d\n", node.data)
	}
}
