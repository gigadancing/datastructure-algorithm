package _search

import (
	"container/heap"
	"github.com/eapache/queue"
	"sort"
)

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
	grap := constructGraph(beginWord, wordList)       // 构造无向图
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
func constructGraph(beginWord string, wordList []string) map[string][]string {
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

// 例2-b. 已知两个单词（分别是起始单词与结束单词），一个单词词典，根据转换规则计算所有的从起始单词到结束单词的最短转换路径。
// 转换规则如下：
// 1. 在转换时，只能转换单词中的一个字符
// 2. 转换得到的新单词必须在单词词典中
// 例如：
// beginWord="hit";endWord="cog";wordList=["hot","dot","dog","lot","log","cog"]
// 最短转化路径为：["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]
func FindLadders(beginWord, endWord string, wordList []string) [][]string {
	ladders := make([][]string, 0) // 结果
	graph := constructGraph2(beginWord, wordList)
	q := make([]*Qitem, 0)
	endWordPos := make([]int, 0)

	BfsGraph(beginWord, endWord, &graph, &q, &endWordPos)

	for _, pos := range endWordPos {
		path := make([]string, 0)
		for pos != -1 { // 从endWord到beginWord将路径上所有节点加入path
			path = append(path, q[pos].Word)
			pos = q[pos].ParentPos
		}

		lad := make([]string, 0)
		for i := len(path) - 1; i >= 0; i-- { // 倒序输出到lad
			lad = append(lad, path[i])
		}
		ladders = append(ladders, lad) // lad加入结果
	}

	return ladders
}

//
type Qitem struct {
	Word      string
	ParentPos int
	Steps     int
}

//
func NewQitem(word string, parentPos int, steps int) *Qitem {
	return &Qitem{
		Word:      word,
		ParentPos: parentPos,
		Steps:     steps,
	}
}

//
func BfsGraph(beginWord, endWord string, graph *map[string][]string, q *[]*Qitem, endWordPos *[]int) {
	visit := make(map[string]int, 0) // <word,steps>
	visit[beginWord] = 1
	*q = append(*q, NewQitem(beginWord, -1, 1)) //起始节点的前驱为-1
	front := 0                                  // 指向队列头
	minSteps := 0                               // 到达endWord的最小步数

	for front != len(*q) { // front超过了队列尾部
		word := (*q)[front].Word
		steps := (*q)[front].Steps

		if minSteps != 0 && steps > minSteps { // steps>minSteps代表所有到达终点的路径都已搜索完
			break
		}

		if word == endWord { // 搜索到结果
			minSteps = steps // 记录到达的最小步数
			*endWordPos = append(*endWordPos, front)
		}

		for _, w := range (*graph)[word] {
			if _, ok := visit[w]; !ok || visit[w] == steps+1 { // 节点未被搜索或另一条路径更短
				*q = append(*q, NewQitem(w, front, steps+1))
				visit[w] = steps + 1
			}
		}

		front++
	}
}

//
func constructGraph2(beginWord string, wordList []string) map[string][]string {
	graph := make(map[string][]string, 0)
	hasBeginWord := false
	for _, w := range wordList { // wordList中可能有beginWord，直接加入可能出现重复，故先判断wordList中是否有beginWord
		if beginWord == w {
			hasBeginWord = true
		}
		graph[w] = make([]string, 0)
	}

	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if connect(wordList[i], wordList[j]) {
				graph[wordList[i]] = append(graph[wordList[i]], wordList[j])
				graph[wordList[j]] = append(graph[wordList[j]], wordList[i])
			}
		}
		if hasBeginWord == false && connect(beginWord, wordList[i]) {
			graph[beginWord] = append(graph[beginWord], wordList[i])
			graph[wordList[i]] = append(graph[wordList[i]], beginWord)
		}
	}

	return graph
}

// 例3. 火柴棍摆正方形
// 已知一个数组，保存了n个（n<=15）火柴棍，问可否使用这n个火柴棍摆成一个正方形？
// [1,1,2,2,2]:true
// [3,3,4,4,4]:false
// [1,1,2,4,3,2,3]:true
// [1,2,3,4,5,6,7,8,9,10,5,4,3,2,1]:false
// 优化与剪枝：
// 1. n个火柴的总和对4取余须为0，否则返回假
// 2. 火柴按照从大到小的顺序排序，先尝试大的减少回溯可能
// 3. 每次放置时，每条边不可放置超过综合的1/4长度的火柴
func MakeSquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}
	// 对nums元素求和，若模4余数不为零，返回假
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	// 从大到小达排序
	data := IntSlice(nums[0:])
	sort.Sort(data)
	bucket := [4]int{}
	return generate(0, nums, sum/4, &bucket)
}

func generate(i int, nums []int, target int, bucket *[4]int) bool {
	if i >= len(nums) { // 所有火柴放完
		return bucket[0] == target && bucket[1] == target && bucket[2] == target && bucket[3] == target
	}
	for j := 0; j < 4; j++ {
		if bucket[j]+nums[i] > target {
			continue
		}
		bucket[j] += nums[i]
		if generate(i+1, nums, target, bucket) {
			return true
		}
		bucket[j] -= nums[i]
	}
	return false
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// 例4. 收集雨水
// 已知一个m*n的二维数组，数组存储正整数，代表一个个单元的高度（立方体），将这些立方体想象成水槽，问如果下雨这些立方体会有多少积水。
// 分析：
// 1. 能积水的地面一定不在四周，积水多少与周围最矮的立方体相关。
// 2. 围住中间积水的边界位置不一定在四周，所以“找出四周（边界）上最低的点”求差不可行。
// 思路：
// 1. 搜索队列使用优先级队列（堆），越低的点优先级越高（最小堆），越优先进行搜索。
// 2. 以矩形四周的点作为起始点进行广度优先搜索（这些点要最初加入队列）。
// 3. 使用一个二维数组对进入队列的点进行标记，之后搜索到该点时，不再加入队列。
// 4. 只要队列不空，取出队头元素进行搜索，按四个方向进行拓展，拓展过程中忽略超出边界和已加入队列的点。
// 5. 当对某点(x,y)进行拓展时（h为(x,y)位置的高度，即heightMap[x][y]）得到的新点为(newx,newy)，高度为heightMap[newx][newy]，如果
//    h > heightMap[newx][newy]：
//        结果 += h - heightMap[newx][newy]
//        将heightMap[newx][newy]赋值为h，即升高该位置的水面
//        将(newx,newy)和heightMap[newx][newy]加入队列
func TrapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	q := make(QueueItemSlice, 0)
	rows := len(heightMap)    // 行数
	cols := len(heightMap[0]) // 列数
	mark := make([][]int, 0)  // 空白地图

	// 初始化地图
	for i := 0; i < rows; i++ {
		c := make([]int, 0)
		for j := 0; j < cols; j++ {
			c = append(c, 0)
		}

		mark = append(mark, c)
	}

	//   |---|----------------|---|
	//   | 1 |  4   3   1   3 | 2 |     1  1  1  1  1  1
	//   |   |----------------|   |
	//   | 3 |  2   1   3   2 | 4 |     1  0  0  0  0  1
	//   |   |                |   |
	//   | 4 |  1   3   1   5 | 3 |     1  0  0  0  0  1
	//   |   |----------------|   |
	//   | 2 |  3   3   4   3 | 1 |     1  1  1  1  1  1
	//   |---|----------------|---|

	// 将地图0列和最后一列标记为1
	for i := 0; i < rows; i++ {
		heap.Push(&q, NewQueueItem(i, 0, heightMap[i][0]))
		mark[i][0] = 1
		heap.Push(&q, NewQueueItem(i, cols-1, heightMap[i][cols-1]))
		mark[i][cols-1] = 1
	}

	// 将地图第0行最后一行的剩余部分标记为1
	for i := 1; i < cols-1; i++ {
		heap.Push(&q, NewQueueItem(0, i, heightMap[0][i]))
		mark[0][i] = 1
		heap.Push(&q, NewQueueItem(rows-1, i, heightMap[rows-1][i]))
		mark[rows-1][i] = 1
	}

	// 方向数组
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	waterVolume := 0 // 积水量

	for q.Len() > 0 {
		elem := heap.Pop(&q).(*QueueItem)
		x := elem.X
		y := elem.Y
		h := elem.Height

		for i := 0; i < 4; i++ { // 向四个方向拓展
			newX := x + dx[i]
			newY := y + dy[i]
			// 当拓展点超出边界或已搜索过，跳过该点
			if newX < 0 || newX >= rows || newY < 0 || newY >= cols || mark[newX][newY] == 1 {
				continue
			}
			if h > heightMap[newX][newY] { // 当前点的高度高于拓展点
				waterVolume += h - heightMap[newX][newY]
				heightMap[newX][newY] = h
			}
			heap.Push(&q, NewQueueItem(newX, newY, heightMap[newX][newY]))
			mark[newX][newY] = 1
		}
	}

	return waterVolume
}

//
type QueueItem struct {
	X, Y   int
	Height int
}

//
func NewQueueItem(x, y, h int) *QueueItem {
	return &QueueItem{
		X:      x,
		Y:      y,
		Height: h,
	}
}

// 优先队列（小顶堆）
type QueueItemSlice []*QueueItem

func (q *QueueItemSlice) Less(i, j int) bool {
	return (*q)[i].Height < (*q)[j].Height
}

func (q *QueueItemSlice) Len() int {
	return len(*q)
}

func (q *QueueItemSlice) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *QueueItemSlice) Push(item interface{}) {
	*q = append(*q, item.(*QueueItem))
}

func (q *QueueItemSlice) Pop() interface{} {
	item := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return item
}
