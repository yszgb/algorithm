package models

import (
	"container/list"
	"strconv"
)

// 二叉树相关操作

// 链式存储的二叉树
// 比链表多一个指针

// 定义二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreOrderTraversal(root *TreeNode) []int {
	// 方法一：迭代法，切片实现栈
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root} // 初始化栈，根节点入栈
	for len(stack) > 0 {
		// 弹出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问该节点
		result = append(result, node.Val)

		// 右子树入栈，再左子树入栈。左子树会先处理
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result

	// 方法二：递归法

	// var result []int
	// // 定义递归函数
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	// 根节点开始
	// 	result = append(result, node.Val)
	// 	// 遍历左子树
	// 	traverse(node.Left)
	// 	// 遍历右子树
	// 	traverse(node.Right)
	// }
	// traverse(root)
	// return result
}

// 中序遍历
func InOrderTraversal(root *TreeNode) []int {
	// 方法一：迭代法，使用双向链表作为栈
	result := []int{}
	stack := list.New() // 使用双向链表作为栈
	current := root
	for current != nil || stack.Len() > 0 {
		// 先遍历左子树
		for current != nil {
			stack.PushBack(current)
			current = current.Left
		}

		// 访问栈顶节点
		if stack.Len() > 0 {
			node := stack.Remove(stack.Back()).(*TreeNode) // 弹出栈顶元素
			result = append(result, node.Val)              // 访问节点
			current = node.Right                           // 继续遍历右子树
		}
	}
	return result

	// // 方法二：递归法

	// var result []int
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	traverse(node.Left)
	// 	result = append(result, node.Val)
	// 	traverse(node.Right)
	// }
	// traverse(root)
	// return result
}

// 定义操作类型，用于区分“访问”还是“处理”
type Command struct {
	Action string // "visit" or "process"
	Node   *TreeNode
}

// 后序遍历
func PostOrderTraversal(root *TreeNode) []int {
	// 方法一：统一迭代法
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []Command{{"visit", root}}
	for len(stack) > 0 {
		// 弹出栈顶元素
		command := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if command.Action == "process" {
			// 如果是处理操作，记录节点值
			result = append(result, command.Node.Val)
		} else if command.Node != nil {
			// 后序遍历的顺序是 左 -> 右 -> 根，入栈顺序相反
			stack = append(stack, Command{"process", command.Node}) // 处理当前节点
			if command.Node.Right != nil {
				stack = append(stack, Command{"visit", command.Node.Right}) // 访问右子树
			}
			if command.Node.Left != nil {
				stack = append(stack, Command{"visit", command.Node.Left}) // 访问左子树
			}
		}
	}
	return result

	// 方法二：递归法

	// var result []int
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	traverse(node.Left)
	// 	traverse(node.Right)
	// 	result = append(result, node.Val)
	// }
	// traverse(root)
	// return result
}

// 层次遍历，广度优先搜素（一层一层，从左到右）
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		// 预分配每一层节点的值
		vals := make([]int, 0, len(curLevel))
		nextLevel := make([]*TreeNode, 0, len(curLevel)*2) // 下一次节点数最多是当前层两倍

		// 遍历当前层节点
		for _, node := range curLevel {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		// 添加当前层的节点值到结果中
		res = append(res, vals)
		// res = append([][]int{vals}, res...) // 自底向上层次遍历。还可以遍历 res 交换位置

		curLevel = nextLevel
	}

	return res
}

// Morris 遍历。修改树的结果，空间复杂度 O(1)
func MorrisInOrder(root *TreeNode) []int {
	result := []int{}
	current := root

	for current != nil {
		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
		} else {
			// 找到当前节点左子树的最右节点
			pre := current.Left
			for pre.Right != nil && pre.Right != current {
				pre = pre.Right
			}

			// 建立临时链接
			if pre.Right == nil {
				pre.Right = current
				current = current.Left
			} else {
				// 访问当前节点
				pre.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}

	return result
}

// 判断两个树是否相同
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// 判断两个树是否存在包含关系
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 递归

	if root == nil {
		return false
	}
	if IsSameTree(root, subRoot) {
		return true
	}
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

// 返回所有根节点到叶子节点路径
func BinaryTreePaths(root *TreeNode) []string {
	// 方法一：递归,回溯
	// 隐式回溯，每次递归调用中 s 是通过值传递的，每次递归都会生成一个新的 s
	// 动态构建字符串时，使用 strings.Builder
	// 构建实例后使用 WriteString, WriteRune 函数

	res := make([]string, 0)
	var travel func(node *TreeNode, s string)

	// 递归遍历函数
	travel = func(node *TreeNode, s string) {
		// 判断当前节点是否为叶子节点
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val) // 构建路径字符串
			res = append(res, v)            // 将路径添加到结果中
			return
		}
		s = s + strconv.Itoa(node.Val) + "->" // 更新路径字符串

		// 递归遍历左右子树
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}

	travel(root, "") // 从根节点开始遍历
	return res       // 返回所有路径

	// 方法二：迭代法

	// stack := []*TreeNode{}     // 栈用于存储节点
	// paths := make([]string, 0) // 存储对应节点的路径
	// res := make([]string, 0)   // 最终结果，存储所有路径

	// // 初始化栈和路径
	// if root != nil {
	// 	stack = append(stack, root)
	// 	paths = append(paths, "")
	// }

	// // 深度优先遍历
	// for len(stack) > 0 {
	// 	l := len(stack)     // 当前栈的大小
	// 	node := stack[l-1]  // 获取栈顶节点
	// 	path := paths[l-1]  // 获取对应的路径
	// 	stack = stack[:l-1] // 弹出栈顶节点
	// 	paths = paths[:l-1] // 弹出对应路径

	// 	// 如果是叶子节点，保存当前路径
	// 	if node.Left == nil && node.Right == nil {
	// 		res = append(res, path+strconv.Itoa(node.Val))
	// 		continue
	// 	}

	// 	// 处理右子节点
	// 	if node.Right != nil {
	// 		stack = append(stack, node.Right)                       // 将右子节点压入栈
	// 		paths = append(paths, path+strconv.Itoa(node.Val)+"->") // 更新路径
	// 	}

	// 	// 处理左子节点
	// 	if node.Left != nil {
	// 		stack = append(stack, node.Left)                        // 将左子节点压入栈
	// 		paths = append(paths, path+strconv.Itoa(node.Val)+"->") // 更新路径
	// 	}
	// }

	// return res // 返回所有路径
}

// 求路径和，返回和等于给定值的路径
func PathSum(root *TreeNode, targetSum int) [][]int {
	// 方法一：递归 + 回溯
	res := [][]int{}
	var backtrack func(node *TreeNode, curPath []int, curSum int)
	backtrack = func(node *TreeNode, curPath []int, curSum int) {
		if node == nil {
			return
		}

		// 添加当前节点到路径中
		curPath = append(curPath, node.Val)
		curSum += node.Val

		// 判断节点是否为叶子节点，路径和是否等于目标值
		if node.Left == nil && node.Right == nil && curSum == targetSum {
			res = append(res, append([]int{}, curPath...))
		} else {
			// 递归，遍历左右子树
			backtrack(node.Left, curPath, curSum)
			backtrack(node.Right, curPath, curSum)
		}

		// 回溯，移除当前节点
		// 这里不需要显式进行回溯，每次递归调用中我们都创建了一个新的切片
		// curPath = curPath[:len(curPath)-1]
	}

	backtrack(root, []int{}, 0)

	return res

	// 方法二：迭代
	// res := [][]int{}
	// if root == nil {
	// 	return res
	// }

	// // 使用栈存储节点、当前路径、路径和
	// stack := []struct { // 创建栈，存放结构体数组
	// 	node *TreeNode
	// 	path []int
	// 	sum  int
	// }{{root, []int{}, 0}} // 先将根节点加入到栈中

	// for len(stack) > 0 {
	// 	// 获取栈顶的结构体元素
	// 	top := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1] // 弹出栈顶

	// 	node := top.node
	// 	curPath := append(top.path, node.Val)
	// 	curSum := top.sum + node.Val

	// 	// 检查节点是否为叶子节点，路径和是否等于目标值
	// 	if node.Left == nil && node.Right == nil && curSum == targetSum {
	// 		// 深拷贝
	// 		res = append(res, append([]int{}, curPath...)) // 先将 curPath 传入空切片，再加入 res。避免后序更改 curPath 对 res 造成影响
	// 	}

	// 	// 右子树、左子树入栈，会优先处理左子树
	// 	if node.Right != nil {
	// 		stack = append(stack, struct {
	// 			node *TreeNode
	// 			path []int
	// 			sum  int
	// 		}{node.Right, curPath, curSum})
	// 	}
	// 	if node.Left != nil {
	// 		stack = append(stack, struct {
	// 			node *TreeNode
	// 			path []int
	// 			sum  int
	// 		}{node.Left, curPath, curSum})
	// 	}
	// }

	// return res
}
