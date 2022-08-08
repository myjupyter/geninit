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
	AA                  *sql.DB
	AExample            int
	BExample            string
	AnExampleOfLongName []struct {
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
	DExample struct {
		A int
	}
	EExample Str
	mExample map[string]interface{}
	IExample any
	FExample chan struct{}
	GExample []struct {
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
