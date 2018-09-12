package main

import "fmt"

type ElementType string

type BiTreeNode struct {
	data   ElementType
	Lchild *BiTreeNode
	Rchild *BiTreeNode
	count  int
}

type BiTree *BiTreeNode

func createBiTree() *BiTreeNode {
	var ch ElementType
	fmt.Println("请输入添加的数据：")
	fmt.Scan(&ch)
	var T BiTree

	if ch == "*" {
		T = nil
	} else {
		T = BiTree(new(BiTreeNode))
		(*T).data = ch
		(*T).Lchild = createBiTree()
		(*T).Rchild = createBiTree()
	}
	return T
}

func preOrderIterate(T BiTree) {
	if T == nil {
		return
	}

	fmt.Println(T.data)
	preOrderIterate(T.Lchild)
	preOrderIterate(T.Rchild)
}

func inOrderIterate(T BiTree) {
	if T == nil {
		return
	}

	inOrderIterate(T.Lchild)
	fmt.Println(T.data)
	inOrderIterate(T.Rchild)
}

func postOrderIterate(T BiTree) {
	if T == nil {
		return
	}

	postOrderIterate(T.Lchild)
	postOrderIterate(T.Rchild)
	fmt.Println(T.data)
}

func NodeNum(T BiTree) int {
	if T == nil {
		return 0
	} else {
		return (1 + NodeNum(T.Lchild) + NodeNum(T.Rchild))
	}
}

func DepthOfTree(T BiTree) int {
	if T == nil {
		return 0
	} else {
		leftDepth := DepthOfTree(T.Lchild)
		rightDepth := DepthOfTree(T.Rchild)
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}

func LeafNum(T BiTree) int {
	if T == nil {
		return 0
	} else if T.Lchild == nil && T.Rchild == nil {
		return 1
	} else {
		return (LeafNum(T.Lchild) + LeafNum(T.Rchild))
	}
}

func findData(T BiTree, data ElementType) BiTree {
	if T == nil {
		return nil
	}

	if T.data == data {
		return T
	} else {
		t := findData(T.Lchild, data)
		if t != nil {
			return t
		}
		return findData(T.Rchild, data)
	}
}

func main() {
	t := createBiTree()
	fmt.Println("前序遍历：")
	preOrderIterate(t)

	fmt.Println("中序遍历：")
	inOrderIterate(t)

	fmt.Println("后序遍历：")
	postOrderIterate(t)

	fmt.Println("二叉树的节点数量：")
	fmt.Println(NodeNum(t))

	fmt.Println("二叉树的深度为：")
	fmt.Println(DepthOfTree(t))

	fmt.Println("二叉树的叶子节点为：")
	fmt.Println(LeafNum(t))

	fmt.Println("查找二叉树")
	dataNode := findData(t, ElementType("b"))
	if dataNode == nil {
		fmt.Println("没有找到节点")
	} else {
		fmt.Println(dataNode)
	}
}
