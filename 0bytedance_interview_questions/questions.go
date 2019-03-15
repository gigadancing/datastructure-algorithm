package biq

import "fmt"

// 有一个n边形，顶点为(P0 P1 ... Pn-1)，每一边都是垂直或水平线段，现给定数值K，以P0为起点将n边形分成k段，每段长度相同，请打印出所有
// 的k等分点(K0 K1 ... Kn-1)的坐标。
const (
	HORIZONTAL = iota
	VERTICAL
)

type Point struct {
	x, y int // 横纵坐标
}

type Edge struct {
	start, end Point // 起点终点
}

func (e *Edge) Length() int {
	return dis(e.start, e.end)
}

func (e *Edge) Direction() int {
	if e.start.x == e.end.x {
		return VERTICAL
	}
	return HORIZONTAL
}

// 两点间距离
func dis(p1, p2 Point) int {
	if p1.x == p2.x {
		return sub(p1.y, p2.y)
	}
	return sub(p1.x, p2.x)
}

//
func sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func FindAllEquivalentPoints(nums []Point, k int) []Point {
	if len(nums) < 3 {
		return nil
	}
	eqPoints := make([]Point, 0)
	eqPoints = append(eqPoints, nums[0]) // 将多边形起点加入结果

	edges := make([]Edge, 0)
	edgesLen := make([]int, 0)
	perimeter := 0
	for i := 0; i+1 <= len(nums)-1; i++ {
		if nums[i].x == nums[i+1].x {
			edges = append(edges, Edge{nums[i], nums[i+1]})
		} else {
			edges = append(edges, Edge{nums[i], nums[i+1]})
		}
		d := dis(nums[i], nums[i+1])

		edgesLen = append(edgesLen, d)
		perimeter += d
	}
	edges = append(edges, Edge{nums[len(nums)-1], nums[0]})
	edgesLen = append(edgesLen, dis(nums[len(nums)-1], nums[0]))
	perimeter += dis(nums[len(nums)-1], nums[0])

	n := perimeter / k // 等分线段长度

	sum := 0
	for i := 0; i < len(nums) && k >= 0; i++ {
		sum += edges[i].Length()
		if sum >= n {
			sum -= n
			k--
			if edges[i].Direction() == HORIZONTAL {
				if edges[i].end.x-sum > 0 {
					eqPoints = append(eqPoints, Point{edges[i].end.x - sum, edges[i].end.y})
				} else {
					eqPoints = append(eqPoints, Point{sum + edges[i].end.x, edges[i].end.y})
				}
			} else {
				if edges[i].end.y-sum > 0 {
					eqPoints = append(eqPoints, Point{edges[i].end.x, edges[i].end.y - sum})
				} else {
					eqPoints = append(eqPoints, Point{edges[i].end.x, edges[i].end.y + sum})
				}
			}
		}
	}

	return eqPoints
}

// 用单向链表表示十进制整数，求两个整数的和。如下图：1234+34=1268，请注意单向链表的方向，不允许使用其他数据结构。
// 1 --> 2 --> 3 --> 4
// +           3 --> 4
// --------------------
// 1 --> 2 --> 6 --> 8
type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		val: val,
	}
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead, ptr *ListNode
	ptr = head
	for head != nil {
		head = head.next
		ptr.next = newHead
		newHead = ptr
		ptr = head
	}
	return newHead
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.next
	}
	fmt.Println()
}

func AddList(head1, head2 *ListNode) *ListNode {
	head1 = Reverse(head1)
	head2 = Reverse(head2)
	ptr1 := head1
	ptr2 := head2

	for ptr1 != nil && ptr2 != nil {
		ptr1.val += ptr2.val
		if ptr1.val >= 10 {
			ptr1.val -= 10
			if ptr1.next != nil {
				ptr1.next.val += 1
			} else {
				node := NewListNode(1)
				ptr1.next = node
			}
		}
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}

	return Reverse(head1)
}

// 有一组不同高度的台阶，由一个整数数组表示，数组中每个数是台阶的高度。当开始下雨了，台阶之间能积多少水。
// 如数组为[0,1,0,2,1,0,1,3,2,1,2,1]
func AccumulatedWater(nums []int) int {
	begin := 1 // 两端不可能积水，从第2个开始
	n := len(nums)
	sum := 0
	for begin < n-1 {
		for i := begin + 1; i < n; i++ {
			if nums[i] >= nums[begin] {
				sum += getWaterVolume(begin, i, nums)
				begin = i
			}
		}
		begin++
	}

	return sum
}

func getWaterVolume(begin, end int, nums []int) int {
	sum := (end - begin - 1) * nums[begin]
	for i := begin + 1; i <= end-1; i++ {
		sum -= nums[i]
	}
	return sum
}
