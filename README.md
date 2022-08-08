# geninit

##### This tool generates options pattern for your struct

---

# Usage 
```
Usage of geninit:
  -filename string
    	select file with structs
  -optprefix string
    	set prefix for Option type
  -property string
    	pass properties for fields (example: -p "A:required;B:alias=BB")
  -type string
    	selecte type which for you want generate initialization
```

Works only with composite types

---

# Install

```
$ go install github.com/myjupyter/geninit@v0.1.1
```

---

# Example 

See examples in example/*.go

```
$ geninit --filename=./example/example.go --type=Str --optprefix="Str" --property="D:required;E:required;F:required;A:alias=LONGALIASFORFUNC" > ./example/str_options.go
```
### example.go
```go
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
```

### Result

```go
package example

import (
  "errors"
  json "encoding/json"
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
```
