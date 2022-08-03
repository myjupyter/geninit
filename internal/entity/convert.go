package entity

import (
	"errors"
	"go/ast"
	"regexp"
	"strings"

	"github.com/fatih/structtag"
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
	if f.Names == nil {
		return nil, errors.New("field does not have name")
	}
	field := &Field{
		Name: f.Names[0].Name,
	}
	if f.Tag != nil && len(f.Tag.Value) != 0 {
		tags, err := structtag.Parse(strings.Trim(f.Tag.Value, "`"))
		if err != nil {
			return nil, err
		}
		tag, err := tags.Get(genInitTagName)
		if err == nil {
			field.TagValue = tag.Value()
		}
	}
	switch t := f.Type.(type) {
	case *ast.Ident:
		field.TypeName = t.Name
	case *ast.StructType, *ast.MapType, *ast.ArrayType, *ast.ChanType, *ast.InterfaceType:
		field.TypeName = prepareType(string(c.RawFile[f.Pos():f.End()]))
	default:
	}
	if err := field.parseTagValue(); err != nil {
		return nil, err
	}
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
