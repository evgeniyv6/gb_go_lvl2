// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package main

import "sync"

type MyIntStack struct {
	lock sync.Mutex
	Items []MyInt
}

func NewMyIntStack() *MyIntStack {
	return &MyIntStack{
		sync.Mutex{},
		[]MyInt{},
	}
}

func (st *MyIntStack) Push(item MyInt) {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.Items = append(st.Items, item)
}

func (st *MyIntStack) Pop() MyInt {
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

func (st *MyIntStack) isEmpty() bool {
	return len(st.Items) == 0
}
