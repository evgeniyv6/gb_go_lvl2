package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"strings"
)

const constTeml = `// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package {{.Package}}
type {{.Name}}s []{{.Name}}
func (c {{.Name}}s) List() []{{.Name}} {
return []{{.Name}}{{"{"}}{{.List}}{{"}"}}
}`


const (
	pkgName = "main"
	srcFileName = "src.go"
	genFileName = "list_gen.go"
	typeName = "Color"
)

func mainCHANGETOMAINONLY() {
	consts, err := getConsts(srcFileName, typeName)

	if err != nil {
		log.Fatal(err)
	}

	templateData:= struct {
		Package string
		Name string
		List string
	}{
		pkgName,
		genFileName,
		strings.Join(consts, ", "),
	}

	genFile, err := os.Create(genFileName)
	if err != nil {
		log.Fatal(err)
	}


	t:= template.Must(template.New("const-list").Parse(constTeml))

	if err := t.Execute(genFile, templateData); err != nil {
		log.Fatal(err)
	}

}

func getConsts(srcFileName, typeName string) ([]string, error){
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, srcFileName, nil, 0)
	if err != nil {
		return nil, err
	}

	var (
		constType string
		out []string
	)

	for _, decl:= range astFile.Decls {
		genDecl, ok:=decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			vspec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if vspec.Type == nil && len(vspec.Values) > 0 {
				constType = ""
				continue
			}

			if vspec.Type != nil {
				ident, ok:= vspec.Type.(*ast.Ident)
				if !ok {
					continue
				}
				constType = ident.Name
			}
			if constType == typeName {
				if len(vspec.Names) == 0{
					continue
				}
				out = append(out, vspec.Names[0].Name)
			}
		}

		}
	return out, nil
}

