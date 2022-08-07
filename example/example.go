package example

import (
	"database/sql"
	json "encoding/json"
	"go/ast"
)

type InterfaceEmpty interface {
}

type Interface interface {
	Do() error
}

type Str struct {
	A int
	B *ast.ArrayType
	C json.RawMessage
	D *map[string]struct{}
	E *chan int
	F interface{}
	G *interface{}
}

type ExampleStruct struct {
	AA *sql.DB
	A  int `geninit:"alias=Struct"`
	B  string
	C  []struct {
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
	}
	D struct {
		A int
	}
	E Str
	m map[string]interface{}
	I any
	F chan struct{}
	G []struct {
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
	}
}

type ExampleStructs []ExampleStruct

func foo() {
	type exampleStruct struct {
		a int
	}
}

func bar() func() func() {
	return func() func() {
		return func() {
			type n map[string]interface{}
			return
		}
	}
}
