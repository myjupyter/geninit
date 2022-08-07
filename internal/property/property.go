package property

import (
	"errors"
	"fmt"
	"strings"
)

type Properties struct {
	OptionType     *OptionType
	FieldPropertis []*FieldProperty
}
type OptionType struct {
	Prefix string
}

type FieldProperty struct {
	Name string

	Required bool
	Alias    string
}

func Parse(px, ps string) (*Properties, error) {
	fps, err := ParseFieldsProperty(ps)
	if err != nil {
		return nil, err
	}
	return &Properties{
		OptionType: &OptionType{
			Prefix: px,
		},
		FieldPropertis: fps,
	}, nil
}

func ParseFieldsProperty(p string) ([]*FieldProperty, error) {
	if p == "" {
		return nil, nil
	}
	var fs []*FieldProperty
	for _, fieldProp := range strings.Split(p, ";") {
		ss := strings.Split(fieldProp, ":")
		if len(ss) != 2 {
			return nil, errors.New(`delimiter ":" has to separate variable name and its properties`)
		}
		f := &FieldProperty{
			Name: ss[0],
		}
		props := strings.Split(ss[1], ",")
		if len(ss) == 0 {
			return nil, errors.New("no properties were passed")
		}
		for _, prop := range props {
			if strings.HasPrefix(prop, "required") {
				f.Required = true
			} else if strings.HasPrefix(prop, "alias=") {
				if f.Alias != "" {
					return nil, fmt.Errorf("alias property was already set for %s field", f.Name)
				}
				if len(prop) == len("alias=") {
					return nil, fmt.Errorf("alias property can't be empty for %s field", f.Name)
				}
				f.Alias = prop[6:]
			} else {
				return nil, fmt.Errorf("unexpected field property %s", prop)
			}
		}
		fs = append(fs, f)
	}
	return fs, nil
}
