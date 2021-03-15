package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

const tmpl = `// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package {{.Package}}

import "sync"

type {{.TypeName}}Stack struct {
	lock sync.Mutex
	Items []{{.TypeName}}
}

func New{{.TypeName}}Stack() *{{.TypeName}}Stack {
	return &{{.TypeName}}Stack{
		sync.Mutex{},
		[]{{.TypeName}}{},
	}
}

func (st *{{.TypeName}}Stack) Push(item {{.TypeName}}) {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.Items = append(st.Items, item)
}

func (st *{{.TypeName}}Stack) Pop() {{.TypeName}} {
	st.lock.Lock()
	defer st.lock.Unlock()

	if len(st.Items) == 0 {
		panic("error tmpl")
	}
	
	l := len(st.Items)
	oldItem := st.Items[l-1]
	st.Items = st.Items[:l-1]
	return oldItem
}

func (st *{{.TypeName}}Stack) isEmpty() bool {
	return len(st.Items) == 0
}
`

func main() {
	t := template.Must(template.New("myStack").Parse(tmpl))
	for i:= 1; i<len(os.Args); i++ {
		dst := strings.ToLower(os.Args[i]) + "_stack.go"
		f, err := os.Create(dst)
		if err != nil {
			fmt.Printf("cannot create %s. Error: %s\n", dst,err)
			continue
		}

		params := map[string]string{
			"TypeName": os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}

		if err := t.Execute(f, params); err != nil {
			log.Fatalf("Cannot exec template %s\n", err)
		}

		if err := f.Close(); err != nil {
			log.Fatalf("Cannot close template file %s\n", err)
		}
	}
}

