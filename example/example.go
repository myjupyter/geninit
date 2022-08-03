package example

type InterfaceEmpty interface {
}

type Interface interface {
	Do() error
}

type Str struct {
	A int
}

type ExampleStruct struct {
	A int `geninit:"alias=Struct"`
	B string
	C []struct {
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
