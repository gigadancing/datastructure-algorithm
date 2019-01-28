package myset

import "sync"

type Set struct {
	m map[int]bool
	sync.RWMutex
}

// 构造函数
func New() *Set {
	return &Set{
		m: map[int]bool{},
	}
}

// 添加元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// 删除元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// 元素是否存在
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 元素个数
func (s *Set) Len() int {
	return len(s.m)
}

// 清空集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

// 集合是否为空
func (s *Set) Empty() bool {
	return s.Len() == 0
}

// 将集合元素转化为数组形式
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	data := make([]int, 0)
	for key := range s.m {
		data = append(data, key)
	}
	return data
}
