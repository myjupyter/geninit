package example

import (
	json "encoding/json"
	"errors"
	"go/ast"
)

type StrOption func(*Str) error

func NewStr(d *map[string]struct{}, e *chan int, f interface{}, opts ...StrOption) (*Str, error) {
	obj := &Str{
		D: d,
		E: e,
		F: f,
	}
	for _, opt := range opts {
		if err := opt(obj); err != nil {
			return nil, err
		}
	}
	return obj, nil
}

func MustStr(d *map[string]struct{}, e *chan int, f interface{}, opts ...StrOption) *Str {
	obj, err := NewStr(d, e, f, opts...)
	if err != nil {
		panic(err)
	}
	return obj
}

func WithLONGALIASFORFUNC(a int) StrOption {
	return func(obj *Str) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithB(b *ast.ArrayType) StrOption {
	return func(obj *Str) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithC(c json.RawMessage) StrOption {
	return func(obj *Str) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithG(g *interface{}) StrOption {
	return func(obj *Str) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}
