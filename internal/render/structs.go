package render

import "strings"

type File struct {
	PackageName string
	NeedNew     bool
	Imports     []*Import
	Struct      *Struct
}

type Import struct {
	Alias string
	Name  string
	Path  string
}

type Struct struct {
	TypeName string
	Prefix   string
	Fields   []*Field
}

func (s Struct) ListRequiredFields() []*Field {
	var ns []*Field
	for _, f := range s.Fields {
		if f.Required {
			ns = append(ns, f)
		}
	}
	return ns
}

func (s Struct) RequiredFieldsAsParams() string {
	var ss []string
	for _, f := range s.Fields {
		if f.Required {
			ss = append(ss, f.ParamShorthand+" "+f.TypeName)
		}
	}
	return strings.Join(ss, ", ")
}

func (s Struct) PassRequiredFields() string {
	var ss []string
	for _, f := range s.Fields {
		if f.Required {
			ss = append(ss, f.ParamShorthand)
		}
	}
	return strings.Join(ss, ", ")
}

func (s Struct) HasRequiredAndOptionalField() bool {
	return s.HasOptionalField() && s.HasRequiredField()
}

func (s Struct) HasRequiredField() bool {
	for _, f := range s.Fields {
		if f.Required {
			return true
		}
	}
	return false

}

func (s Struct) HasOptionalField() bool {
	for _, f := range s.Fields {
		if !f.Required {
			return true
		}
	}
	return false
}

type Field struct {
	Name           string
	TypeName       string
	Alias          string
	Required       bool
	ParamShorthand string
}

func (f Field) WithName() string {
	if f.Alias != "" {
		return f.Alias
	}
	return f.Name
}
