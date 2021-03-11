package iqueue

import (
	"reflect"
	"sort"
	"sync"
)

type Item interface {}

type Queue struct {
	Items []Item
	lock sync.Mutex
}

// если убрать lock, получим состояние гонки (задание 3)
func (q *Queue) Enqueue(elem Item) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Items = append(q.Items, elem)
}

func (q *Queue) Dequeue() Item {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.Items) == 0 {
		return nil
	}

	oldItem := q.Items[0]
	q.Items = q.Items[1:]
	return oldItem
}

func (q *Queue) Sort() {
	q.lock.Lock()
	defer q.lock.Unlock()

	sort.Slice(q.Items, func(i, j int) bool {
		if reflect.TypeOf(q.Items[i]).String() == "int" &&
			reflect.TypeOf(q.Items[j]).String() == "int" {
			return q.Items[i].(int) < q.Items[j].(int)
		}

		return false
	})
}