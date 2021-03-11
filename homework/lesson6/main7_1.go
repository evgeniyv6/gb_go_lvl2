package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"
)

const count = 10

type Item int

type Queue struct {
	Items []Item
	lock sync.Mutex
}

func (q *Queue) Enqueue(elem Item) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Items = append(q.Items, elem)
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.Items) == 0
}

func (q *Queue) Dequeue() Item {
	//q.lock.Lock()
	//defer q.lock.Unlock()

	if len(q.Items) == 0 {
		return -1
	}
	oldItem := q.Items[0]
	q.Items = q.Items[1:]
	return oldItem
}

func (q *Queue) Sort() {
	q.lock.Lock()
	defer q.lock.Unlock()
	sort.Slice(q.Items, func(i,j int) bool {
		return q.Items[i] < q.Items[j]
	})
}

func main2() {
	var (
		q Queue
		wg sync.WaitGroup
	)
	wg.Add(count)
	for i:=0;i<count;i++ {
		go func(i int) {
			defer wg.Done()
			q.Enqueue(Item(i));
			time.Sleep(1*time.Second)
		}(i)
	}
	for i:=0;i<count;i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			q.Dequeue();
		}(i)
	}


	wg.Wait()
	q.Sort()
	fmt.Println(q.Items)
}

func main() {
	main2()
	//infiniteStack()
}

func infiniteStack() {
	wg:=sync.WaitGroup{}
	for i:=0;i<count;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(5*time.Minute)
		}()
	}
	runtime.Gosched()
	wg.Wait()
}