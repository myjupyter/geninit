package entity

import (
	"errors"
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

const (
	openClauseExpr  = `{\n(\t)+`
	closeClauseExpr = `\n(\t)+}`
	identExpr       = `\n(\t)+`
)

var (
	openClauseRegex  = regexp.MustCompile(openClauseExpr)
	closeClauseRegex = regexp.MustCompile(closeClauseExpr)
	identRegex       = regexp.MustCompile(identExpr)
)

type convertor struct {
	RawFile []byte
}

func (c convertor) convertToField(f *ast.Field) (*Field, error) {
	if len(f.Names) == 0 {
		return nil, errors.New("field does not have name")
	}
	field := &Field{
		Name: f.Names[0].Name,
	}
	switch t := f.Type.(type) {
	case *ast.Ident:
		field.TypeName = t.Name
	case *ast.StructType, *ast.MapType, *ast.ArrayType, *ast.ChanType, *ast.InterfaceType, *ast.StarExpr, *ast.SelectorExpr:
		trimName := f.Names[0].Name
		typeName := strings.TrimPrefix(string(c.RawFile[f.Pos()-1:f.End()]), trimName)
		field.TypeName = prepareType(typeName)
	default:
		return nil, fmt.Errorf("unexpected field type: %T", f.Type)
	}
	field.ImportPackage = getPackageImport(field.TypeName)
	return field, nil
}

func (c convertor) convertToStruct(s *ast.StructType) (*Struct, error) {
	if s.Fields == nil || len(s.Fields.List) == 0 {
		return nil, nil
	}
	st := &Struct{}
	for _, field := range s.Fields.List {
		f, err := c.convertToField(field)
		if err != nil {
			return nil, err
		}
		st.Fields = append(st.Fields, f)
	}
	return st, nil

}

func prepareType(s string) string {
	s = strings.TrimSpace(s)
	s = openClauseRegex.ReplaceAllString(s, "{")
	s = closeClauseRegex.ReplaceAllString(s, "}")
	s = identRegex.ReplaceAllString(s, "; ")
	return s
}

func getPackageImport(tn string) string {
	idx := strings.Index(tn, ".")
	if idx == -1 {
		return ""
	}
	pkg := tn[:idx]
	if pkg[0] == '*' {
		return pkg[1:]
	}
	return pkg
}
