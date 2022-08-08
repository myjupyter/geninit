package example

import (
	"database/sql"
	"errors"
)

type Option func(*ExampleStruct) error

func NewExampleStruct(opts ...Option) (*ExampleStruct, error) {
	obj := &ExampleStruct{}
	for _, opt := range opts {
		if err := opt(obj); err != nil {
			return nil, err
		}
	}
	return obj, nil
}

func MustExampleStruct(opts ...Option) *ExampleStruct {
	obj, err := NewExampleStruct(opts...)
	if err != nil {
		panic(err)
	}
	return obj
}

func WithAA(aa *sql.DB) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithAExample(ae int) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithBExample(be string) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithAnExampleOfLongName(aeoln []struct {
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
}) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithDExample(de struct{ A int }) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithEExample(ee Str) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithmExample(me map[string]interface{}) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithIExample(ie any) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithFExample(fe chan struct{}) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}

func WithGExample(ge []struct {
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
}) Option {
	return func(obj *ExampleStruct) error {
		// TODO: your initialization rule
		return errors.New("not implemented")
	}
}
