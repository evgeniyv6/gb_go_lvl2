package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

const (
	count = 10
)

func main() {
	var tt = []int{}
	for i:=0;i<10;i++ {
		tt = append(tt,i)
	}
	fmt.Println("tt =",tt[:4])
	//nThreads()
	// tryDefereMutex()
	st := Stack{}
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i:=0; i<count; i++ {
		go func(i int) {
			st.Push(i)
			wg.Done()
		}(i)
	}
	// st.Pop()
	wg.Wait()
	sort.Slice(st.Items, func(i,j int) bool {
		return st.Items[i].(int) < st.Items[j].(int)
	})
	fmt.Println(st.Items)


	res, err := muliplyMatrices(); if err != nil {
		panic(err)
	}
	fmt.Println(res)


}

func criticalSection (m2 *sync.Mutex) {
	defer m2.Unlock ()
}

func nThreads() {
	var (
		wg = sync.WaitGroup{}
	)
	wg.Add(count)
	for i:=0; i<count; i++ {
		go func() {
			fmt.Println("goro #", i)
			time.Sleep(10 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("All %d goroutins completed\n", count)
}

type Item interface{}

type Stack struct {
	lock sync.Mutex
	Items []Item
}

func (s *Stack) Push(elem Item) {
	s.lock.Lock()
	muUnlock(&s.lock)
	// defer s.lock.Unlock()
	s.Items = append(s.Items, elem)

}

func (s *Stack) Pop() Item{
	s.lock.Lock()
	muUnlock(&s.lock)

	l := len(s.Items)
	if l == 0 {
		return nil
	}
	old := s.Items[l-1]
	s.Items = s.Items[:l-1]
	return old
}

func (s *Stack) isEmpty() bool {
	s.lock.Lock()
	muUnlock(&s.lock)
	return len(s.Items) == 0
}

func tryDefereMutex() {
	var m sync.Mutex
	m.Lock()
	muUnlock(&m)
}

func muUnlock(m *sync.Mutex) {
	defer m.Unlock()
}

func muliplyMatrices() ([][]int, error) {
	m1 := [][]int{
		{1,1,1},
		{2,2,2},
		{3,3,3},
	}
	m2 := [][]int{
		{4,4,4},
		{5,5,5},
		{6,6,6},
	}

	res := make([][]int, len(m1))

	wg := sync.WaitGroup{}
	m1ColNum := len(m1[0])
	m2RowNum := len(m2)
	if m1ColNum != m2RowNum {
		return nil, errors.New("matrices cannot be multiplied")
	}

	for i:=0; i< len(m1); i++ {
		res[i] = make([]int, len(m2[0]))
		for j:=0; j < len(m2[0]); j++ {
			for k:=0; k < len(m2); k++ {
				wg.Add(1)
				go func(i,j,k int) {
					res[i][j] += m1[i][k] * m2[k][j]
					wg.Done()
				}(i,j,k)
			}
		}
	}
	wg.Wait()
	return res, nil
}

type Set struct {
	sync.Mutex
	mm map[int]int
}

func NewSet() *Set {
	return &Set{
		mm: map[int]int{},
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	defer s.Unlock()
	s.mm[i] = i
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type SetRW struct {
	sync.RWMutex
	mm map[int]int
}

func NewSetRW() *SetRW {
	return &SetRW{
		mm: map[int]int{},
	}
}

func (s *SetRW) AddRW(i int) {
	s.Lock()
	defer s.Unlock()
	s.mm[i] = i
}

func (s *SetRW) HasRW(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}