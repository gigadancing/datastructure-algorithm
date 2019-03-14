package biq

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
