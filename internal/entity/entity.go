package entity

import (
	"errors"
	"go/ast"
	"strings"
)

const genInitTagName = "geninit"

type File struct {
	PackageName string
	Structs     []*Struct

	RawFile []byte
}

func (f *File) AddPackageName(name string) error {
	if name == "" {
		return errors.New("package name cannot be empty")
	}
	f.PackageName = name
	return nil
}

func (f *File) AddStructType(s *ast.StructType, name string) error {
	c := convertor{
		RawFile: f.RawFile,
	}
	st, err := c.convertToStruct(s)
	if err != nil {
		return err
	}
	if st == nil {
		return nil
	}
	st.TypeName = name
	f.Structs = append(f.Structs, st)
	return nil
}

type Struct struct {
	TypeName string
	Fields   []*Field
}

type Field struct {
	Name     string
	TypeName string
	TagValue string

	alias   string
	require bool
}

func (f Field) WithName() string {
	if f.alias == "" {
		return f.Name
	}
	return f.alias
}

func (f *Field) parseTagValue() error {
	const (
		requirePrefix = "require"
		aliasPrefix   = "alias="
	)
	exprs := strings.Split(f.TagValue, ",")
	for _, expr := range exprs {
		if strings.EqualFold(expr, requirePrefix) {
			f.require = true
		}
		if strings.HasPrefix(expr, aliasPrefix) {
			// TODO: validate
			f.alias = strings.TrimPrefix(expr, aliasPrefix)
		}
	}
	return nil
}
