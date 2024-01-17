package internal

import (
	"fmt"
	"math"
)

type Node struct {
	Val      int
	Children []*Node
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func reverse(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[len(a)-1-i]
	}
	return res
}

func Task(packSizes []int, target int) []int {
	l := target + target
	for i := range packSizes {
		l = max(l, packSizes[i]+1)
	}

	d := make([]int, l)
	arr := make([][]int, l)
	for i := range d {
		d[i] = -1
	}

	packSizes = reverse(packSizes)

	for _, x := range packSizes {
		d[x] = 1
		arr[x] = []int{x}
		for i := 0; i < l; i++ {
			if d[i] != -1 {
				value := d[i] + 1
				if i+x < l {
					oldValue := l
					if d[i+x] == -1 {
						d[i+x] = min(value, oldValue)
						arr[i+x] = append([]int{}, arr[i]...)
						arr[i+x] = append(arr[i+x], x)
					} else {
						oldValue = d[i+x]
						if value < oldValue {
							d[i+x] = value
							arr[i+x] = append([]int{}, arr[i]...)
							arr[i+x] = append(arr[i+x], x)
						}
					}
				}
			}
		}
	}

	ans := []int{}
	for i := 0; i < l; i++ {
		//fmt.Println(i, d[i], arr[i])
		if i >= target && d[i] != -1 {
			//fmt.Println("got you", arr[i])
			ans = arr[i]
			break
		}
	}
	//fmt.Println(d[target])
	//fmt.Println(arr[target])
	return ans
}

//func Task(packSizes []int, target int) []int {
//
//	root := buildTree(&Node{}, packSizes, target, 0)
//
//	printDFS("", root)
//
//	d, s, ans := findMinDFS(root, math.MaxInt, math.MaxInt, 0, 0, []int{})
//	fmt.Println(d, s, ans)
//
//	return ans[1:]
//}

func printDFS(prefix string, root *Node) {
	if root == nil {
		return
	}
	fmt.Println(prefix, root.Val)
	for i := range root.Children {
		printDFS(prefix+"-", root.Children[i])
	}
}

func findMinDFS(root *Node, minDepth int, minSum int, depth int, sum int, ans []int) (int, int, []int) {
	if root == nil {
		//fmt.Println("got to end", ans, minDepth, depth, sum, minSum)
		if sum < minSum {
			minSum = sum
			minDepth = depth
		}
		if sum == minSum {
			if depth <= minDepth {
				minDepth = depth
			}
		}
		return minDepth, minSum, ans
	}
	lMinDepth := math.MaxInt
	lMinSum := math.MaxInt
	lans := []int{}
	for i := range root.Children {
		d, s, arr := findMinDFS(root.Children[i], minDepth, minSum, depth+1, sum+root.Val, append(ans, root.Val))
		if s < lMinSum {
			lMinSum = s
			lMinDepth = d
			lans = arr
		}
		if s == lMinSum {
			if d < lMinDepth {
				lMinDepth = d
				lans = arr
			}
		}
	}
	return lMinDepth, lMinSum, lans
}

func buildTree(root *Node, packSizes []int, target int, currentSum int) *Node {
	fmt.Println(currentSum)
	if target <= currentSum {
		return nil
	}

	currentSum += root.Val

	for i := range packSizes {
		ch := &Node{
			Val: packSizes[i],
		}
		root.Children = append(root.Children, buildTree(ch, packSizes, target, currentSum))
	}
	return root
}
