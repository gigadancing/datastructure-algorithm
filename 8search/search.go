package _search

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
	(*mark)[y][x] = 1
	// 方向数组
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

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
