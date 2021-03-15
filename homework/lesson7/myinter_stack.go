// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package main

import "sync"

type MyInterStack struct {
	lock sync.Mutex
	Items []MyInter
}

func NewMyInterStack() *MyInterStack {
	return &MyInterStack{
		sync.Mutex{},
		[]MyInter{},
	}
}

func (st *MyInterStack) Push(item MyInter) {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.Items = append(st.Items, item)
}

func (st *MyInterStack) Pop() MyInter {
	st.lock.Lock()
	defer st.lock.Unlock()

	if st.isEmpty() {
		panic("error tmpl")
	}
	
	l := len(st.Items)
	oldItem := st.Items[l-1]
	st.Items = st.Items[:l-1]
	return oldItem
}

func (st *MyInterStack) isEmpty() bool {
	return len(st.Items) == 0
}
