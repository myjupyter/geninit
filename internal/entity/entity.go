package entity

import (
	"errors"
	"go/ast"
	"path"
	"strings"
)

func NewFile(raw []byte) (*File, error) {
	f := &File{
		RawFile: raw,
	}
	return f, nil
}

type File struct {
	PackageName string
	Imports     []*Import
	Structs     []*Struct

	RawFile []byte
}

func (f *File) AddImport(is *ast.ImportSpec) error {
	im := &Import{
		Path: strings.Trim(is.Path.Value, `"`),
	}
	if is.Name != nil {
		im.Alias = is.Name.Name
	}
	im.Name = path.Base(im.Path)
	f.Imports = append(f.Imports, im)
	return nil
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

type Import struct {
	Alias string
	Name  string
	Path  string
}

type Struct struct {
	TypeName string
	Fields   []*Field
}

type Field struct {
	Name          string
	TypeName      string
	ImportPackage string
}
