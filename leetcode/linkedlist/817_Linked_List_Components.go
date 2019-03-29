package linkedlist

// 817.Linked List Components
// We are given head, the head node of a linked list containing unique integer values.
// We are also given the list G, a subset of the values in the linked list.
// Return the number of connected components in G, where two values are connected if they appear consecutively in the
// linked list.
//
// Example 1:
//     Input:
//     head: 0->1->2->3
//     G = [0, 1, 3]
//     Output: 2
// Explanation:
//     0 and 1 are connected, so [0, 1] and [3] are the two connected components.
//
// Example 2:
//     Input:
//     head: 0->1->2->3->4
//     G = [0, 3, 1, 4]
//     Output: 2
// Explanation:
//     0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
// Note:
//    If N is the length of the linked list given by head, 1 <= N <= 10000.
//    The value of each node in the linked list will be in the range [0, N - 1].
//    1 <= G.length <= 10000.
//    G is a subset of all values in the linked list.
// 1.第一种解法：图的深搜
func numComponents(head *ListNode, G []int) int {
	m := make(map[int]int) // 将G中的数据存到map中以便快速访问
	for _, g := range G {
		m[g] = 1
	}
	graph := make(map[int][]int) // 用邻接表表示图
	first := head.Val
	second := 0
	for head.Next != nil { // 遍历链表构造图
		head = head.Next
		second = head.Val
		_, ok1 := m[first]
		_, ok2 := m[second]
		if ok1 && ok2 {
			graph[first] = append(graph[first], second)
			graph[second] = append(graph[second], first)
		}
		first = second
	}
	ans := 0
	visited := make(map[int]int)
	for _, g := range G {
		if _, ok := visited[g]; ok {
			continue
		}
		ans++
		dfs(g, graph, visited)
	}

	return ans
}

// 深搜
func dfs(cur int, graph map[int][]int, visited map[int]int) {
	if _, ok := visited[cur]; ok {
		return
	}
	visited[cur] = 1
	for _, n := range graph[cur] {
		dfs(n, graph, visited)
	}
}

// 1.第二种解法：利用链表的性质
// 遍历链表，若集合节点在集合G中，就标记为1，G=[0, 3, 1, 4]，0->1->2->3->4
// 故最后链表记为1->1->0->1->1，就只需逐个遍历，直到0或nil，就找到一个。
func numComponents2(head *ListNode, G []int) int {
	m := make(map[int]int)
	for _, n := range G {
		m[n] = 1
	}
	ans := 0
	for head != nil {
		_, ok := m[head.Val]
		if head.Next == nil {
			if ok {
				ans++
			}
		} else {
			if _, has := m[head.Next.Val]; !has && ok {
				ans++
			}
		}
		head = head.Next
	}
	return ans
}

// leetcode牛逼代码
func numComponents3(head *ListNode, G []int) int {
	keys := make(map[int]struct{}, len(G))
	for _, i := range G {
		keys[i] = struct{}{}
	}
	count := 0
	flag := false
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := keys[cur.Val]; !ok {
			if !flag && cur != head {
				count++
			}
			// 第一个不存在，count不加，但标志要设置
			flag = true
		} else {
			flag = false
		}
	}
	//结尾那个连续的要删除
	if flag {
		return count
	}
	return count + 1
}
