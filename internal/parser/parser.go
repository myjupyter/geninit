package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type EntityAdder interface {
	AddPackageName(string) error
	AddStructType(*ast.StructType, string) error
}

func ParseTo(filename string, ea EntityAdder) error {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	if err := ea.AddPackageName(f.Name.Name); err != nil {
		return err
	}
	for _, decl := range f.Decls {
		if err := inspectDecl(decl, ea); err != nil {
			return err
		}
	}
	return nil
}

func inspectDecl(n ast.Node, ea EntityAdder) error {
	var err error
	ast.Inspect(n, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			err = inspectTypeSpec(x.Specs, ea)
			if err != nil {
				return false
			}
		default:
			return false
		}
		return true
	})
	return err
}

func inspectTypeSpec(specs []ast.Spec, ea EntityAdder) error {
	for _, spec := range specs {
		switch x := spec.(type) {
		case *ast.TypeSpec:
			return inspectStructType(x, ea)
		default:
		}
	}
	return nil
}

func inspectStructType(ts *ast.TypeSpec, ea EntityAdder) error {
	switch x := ts.Type.(type) {
	case *ast.StructType:
		return ea.AddStructType(x, ts.Name.Name)
	default:
		return nil
	}
}
