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

// 例2-a. 词语阶梯
// 已知两个单词（分别是起始单词和结束单词），一个单词词典，根据转换规则计算从起始单词到结束单词的最短转换步数。
// 转换规则如下：
// 1. 在转换时，只能转换单词中的1个字符
// 2. 转换得到的新单词，必须在单词词典中
// 例如：beginWord="hit",endWord="cog",wordList=["hot","dot","dog","lot","log","cog"]
// 最短转换方式："hit"->"hot"->"dot"->"dog"->"cog"，结果为5
// 思考：
// 单词与单词之间的转换，可以理解为一张图，图的顶点为单词，若两担此之间可以互相转换，则这两个单词所代表的顶点间有一条边。即求图中从一个顶
// 点到另一个顶点的所有路径中，最少包含多少个节点，即为广度优先搜索。
//              dot --- dog
//            /  |       |  \
// hit --- hot   |       |   cog
//            \  |       |  /
//              lot --- log
// 思路：
// 给定图的起始顶点beginWord，终点endWord，图graph，从beginWord开始广度优先搜索图graph，搜索过程中记录到达步数。
// 1. 设置搜索队列Q，队列节点为Pair<顶点，步数>；设置集合visit，记录搜索过的顶点；将<beginWord,1>添加至队列。
// 2. 只要队列不空，取出队列的头元素
//    (1) 若去出的队列头元素为endWord，返回达到前节点的步数
//    (2) 否则拓展该节点，将该节点相邻的且未在visit中的节点与步数同时添加至队列Q，并将所拓展的节点加入visit
// 3. 做最终无法搜索到endWord返回0
func LadderLength(beginWord, endWord string, wordList []string) int {
	grap := constructGrap(beginWord, wordList)        // 构造无向图
	q := queue.New()                                  // 搜索队列
	q.Add(NewGraphNodePair(beginWord, 1))             // 将起始节点加入搜索队列
	visit := make(map[string]*GraphNodePair, 0)       // visit记录已经搜索过的节点
	visit[beginWord] = NewGraphNodePair(beginWord, 1) // 将起始节点加入visit

	for q.Length() > 0 {
		head := q.Peek().(*GraphNodePair)
		steps := head.Steps
		q.Remove()
		if head.Word == endWord {
			return steps
		}
		for _, w := range grap[head.Word] {
			if _, ok := visit[w]; !ok {
				q.Add(NewGraphNodePair(w, steps+1))
				visit[w] = NewGraphNodePair(w, steps+1)
			}
		}
	}

	return 0
}

//
type GraphNodePair struct {
	Word  string
	Steps int
}

//
func NewGraphNodePair(word string, steps int) *GraphNodePair {
	return &GraphNodePair{
		Word:  word,
		Steps: steps,
	}
}

//
func constructGrap(beginWord string, wordList []string) map[string][]string {
	graph := make(map[string][]string, 0)
	wordList = append(wordList, beginWord)
	for _, w := range wordList {
		graph[w] = make([]string, 0)
	}
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if connect(wordList[i], wordList[j]) { // 只有一个字符不同，可以连接
				// 由于是无向图，两个方向都要连接
				graph[wordList[i]] = append(graph[wordList[i]], wordList[j])
				graph[wordList[j]] = append(graph[wordList[j]], wordList[i])
			}
		}
	}

	return graph
}

//
func connect(word1, word2 string) bool {
	cnt := 0 // 记录word1与word2不相等字符的个数
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			cnt++
		}
	}
	return cnt == 1 // 只有一个字符不同才能连接
}
