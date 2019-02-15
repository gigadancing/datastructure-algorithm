package _search

import "github.com/eapache/queue"

// 例1. 岛屿数量
// 用一个二维数组代表一张地图，这张地图由字符'0'与字符'1'组成，其中'0'字符代表水域，'1'字符代表小岛土地，小岛'1'被水'0'所包围，当小岛
// '1'在水平和垂直方向相连接时，认为时同一块土地。求这张地图中小岛的数量。
//  1个小岛       3个小岛
// 1 1 1 1 0    1 1 1 0 0
// 1 1 0 1 0    1 1 0 0 0
// 1 1 0 0 0    0 0 1 0 0
// 0 0 0 0 0    0 0 0 1 1
// 0 0 0 0 0    0 0 0 0 0
func NumIslandsDFS(grid [][]int) [][]int {
	mark := make([][]int, 0)
	for _, row := range grid { // 根据grid构造同样的大小地图，全部标记为0
		r := make([]int, 0)
		for i := 0; i < len(row); i++ {
			r = append(r, 0)
		}
		mark = append(mark, r)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				DFS(&mark, grid, j, i)
			}
		}
	}

	return mark
}

// DFS
// 1. 标记当前搜索位置已被搜索（标记当前位置的mark数组为1）
// 2. 按照方向数组的4个方向，拓展4个新位置newx、newy
// 3. 若新位置不在地图范围内，则忽略
// 4. 如果新位置未曾达到过（mark[newx][newy]为0）且是陆地（grid[newx][newy]为1），继续DFS该位置
//   -------------------------------->x
//   |
//   |              (x, y-1)
//   |
//   |   (x-1, y)   (x,   y)   (x+1, y)
//   |
//   |              (x, y+1)
//   |
//   | y
func DFS(mark *[][]int, grid [][]int, x, y int) {
	// 方向数组
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}
	(*mark)[y][x] = 1
	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		if newY < 0 || newY >= len(*mark) || newX < 0 || newX >= len((*mark)[newY]) { // 超过地图边界
			continue
		}

		if (*mark)[newY][newX] == 0 && grid[newY][newX] == 1 { // 新位置未被探索且新位置为1时，继续搜索
			DFS(mark, grid, newX, newY)
		}
	}
}

// BFS
// 1. 设置搜多队列Q，标记mark[x][y]=1，并将待搜索的位置(x,y)push进入队列Q
// 2. 只要队列不空，取队头元素，按照方向数组的4个方向，拓展4个新位置newx、newy
// 3. 若新位置不在地图内，则忽略
// 4. 若新位置未搜索过且是陆地，将该新位置push进队列，并标记mark[x][y]=1
// ------------------------------------->y
// |
// |                (x-1, y)
// |
// |    (x, y-1)    (x,   y)    (x, y+1)
// |
// |                (x+1, y)
// |
// |
// | x
type Pair struct {
	X, Y int
}

func NewPair(x, y int) *Pair {
	return &Pair{
		X: x,
		Y: y,
	}
}

func NumIslandsBFS(grid [][]int) [][]int {
	mark := make([][]int, 0)
	for _, row := range grid {
		r := make([]int, 0)
		for i := 0; i < len(row); i++ {
			r = append(r, 0)
		}
		mark = append(mark, r)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				BFS(&mark, grid, i, j)
			}
		}
	}
	return mark
}

func BFS(mark *[][]int, grid [][]int, x, y int) {
	// 方向数组
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	(*mark)[x][y] = 1 // 当前位置标记为已搜索
	q := queue.New()
	q.Add(NewPair(x, y)) // 当前位置加入队列

	for q.Length() > 0 {
		x = q.Peek().(*Pair).X
		y = q.Peek().(*Pair).Y
		q.Remove() // 弹出队头元素
		for i := 0; i < 4; i++ {
			// 新位置坐标
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[newX]) { // 超越地图边界
				continue
			}

			if (*mark)[newX][newY] == 0 && grid[newX][newY] == 1 { // 新位置未搜索且为陆地
				q.Add(NewPair(newX, newY)) // 加入队列
				(*mark)[newX][newY] = 1    // 标记为已搜索
			}
		}
	}
}
