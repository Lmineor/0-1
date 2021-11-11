package set

import (
	"sort"
	"sync"
)

type Set struct {
	sync.RWMutex
	m map[string]bool
}

// New 新建集合对象
func New(items ...string) *Set {
	s := &Set{m: make(map[string]bool, len(items))}
	s.Add(items...)
	return s
}

// Add 添加元素
func (s *Set) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

// Remove 删除元素
func (s *Set) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		delete(s.m, item)
	}
}

// In 判断元素是否在集合中
func (s *Set) HasItem(items ...string) bool {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Count() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

// 空集合判断
func (s *Set) Empty() bool {
	return len(s.m) == 0
}

// 无序列表
func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 排序列表
func (s *Set) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}

// Intersect is to get the common items between slices
func (s *Set) Intersect(sets ...*Set) *Set {
	refer := New(s.List()...)
	for _, set := range sets {
		for e := range s.m {
			if _, ok := set.m[e]; !ok {
				delete(refer.m, e)
			}
		}
	}
	return refer
}

// Minux is to get the different items between slices
func (s *Set) Minus(sets ...*Set) *Set {
	refer := s.Union(sets...)
	for _, set := range sets {
		for e := range set.m {
			if _, ok := s.m[e]; ok {
				delete(refer.m, e)
			}
		}
	}
	return refer
}

// Union is to get the all items between slices
func (s *Set) Union(sets ...*Set) *Set {
	rerfer := New(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			rerfer.m[e] = true
		}
	}
	return rerfer
}

// 补集
func (s *Set) Complement(full *Set) *Set {
	rerfer := New()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			rerfer.Add(e)
		}
	}
	return rerfer
}
