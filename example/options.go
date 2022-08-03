package example

import (
	"errors"
)

type StrOption func(*Str) error

func NewStr(opts ...StrOption) (*Str, error) {
	obj := &Str{}
	for _, opt := range opts {
		if err := opt(obj); err != nil {
			return nil, err
		}
	}
	return obj, nil
}

func MustStr(opts ...StrOption) *Str {
	obj, err := NewStr(opts...)
	if err != nil {
		panic(err)
	}
	return obj
}

func WithA(A int) StrOption {
	return func(obj *Str) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

type ExampleStructOption func(*ExampleStruct) error

func NewExampleStruct(opts ...ExampleStructOption) (*ExampleStruct, error) {
	obj := &ExampleStruct{}
	for _, opt := range opts {
		if err := opt(obj); err != nil {
			return nil, err
		}
	}
	return obj, nil
}

func MustExampleStruct(opts ...ExampleStructOption) *ExampleStruct {
	obj, err := NewExampleStruct(opts...)
	if err != nil {
		panic(err)
	}
	return obj
}

func WithStruct(A int) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithB(B string) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithC(C []struct {
	A string
	B int
	C []struct {
		A string
		B string
		i interface {
			Do([]struct {
				A string
				B map[int]struct {
					A int
					B string
				}
			})
		}
	}
}) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithD(D struct{ A int }) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithE(E Str) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func Withm(m map[string]interface{}) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithI(I any) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithF(F chan struct{}) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithG(G []struct {
	A string
	B int
	C []struct {
		A string
		B string
		i interface {
			Do([]struct {
				A string
				B map[int]struct {
					A int
					B string
				}
			})
		}
	}
}) ExampleStructOption {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}
