package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/myjupyter/geninit/internal/entity"
	"github.com/myjupyter/geninit/internal/parser"
	"github.com/myjupyter/geninit/internal/property"
	"github.com/myjupyter/geninit/internal/render"
)

var (
	filename   string
	typename   string
	prefix     string
	properties string
)

func parseFlags() struct{} {
	flag.StringVar(&filename, "filename", "", "select file with structs")
	flag.StringVar(&typename, "type", "", "selecte type which for you want generate initialization")
	flag.StringVar(&prefix, "optprefix", "", "set prefix for Option type")
	flag.StringVar(&properties, "property", "", "pass properties for fields (example: -p \"A:required;B:alias=BB\")")
	flag.Parse()
	return struct{}{}
}

var _ = parseFlags()

func getType(ss []*entity.Struct, t string) (*entity.Struct, error) {
	for _, s := range ss {
		if s.TypeName == t {
			return s, nil
		}
	}
	return nil, fmt.Errorf("no type detected")
}

func genParamShorthand(p string) string {
	var rs []rune
	for i, r := range p {
		if i == 0 && unicode.IsLower(r) {
			rs = append(rs, r)
		}
		if unicode.IsUpper(r) {
			rs = append(rs, unicode.ToLower(r))
		}
	}
	return string(rs)
}

func convertForRendering(f *entity.File, ps *property.Properties) (*render.File, error) {
	s, err := getType(f.Structs, typename)
	if err != nil {
		return nil, err
	}
	matchProp := make(map[string]*property.FieldProperty)
	for _, prop := range ps.FieldPropertis {
		matchProp[prop.Name] = prop
	}
	rs := &render.Struct{
		TypeName: s.TypeName,
		Prefix:   ps.OptionType.Prefix,
	}
	for _, field := range s.Fields {
		prop, ok := matchProp[field.Name]
		rf := &render.Field{
			Name:           field.Name,
			TypeName:       field.TypeName,
			ParamShorthand: genParamShorthand(field.Name),
		}
		if ok {
			rf.Alias = prop.Alias
			rf.Required = prop.Required
			rf.ParamShorthand = genParamShorthand(field.Name)
		}
		rs.Fields = append(rs.Fields, rf)
	}
	matchFields := make(map[string]*entity.Field)
	for _, field := range s.Fields {
		if field.ImportPackage == "" {
			continue
		}
		matchFields[field.ImportPackage] = field
	}
	var ims []*render.Import
	for _, im := range f.Imports {
		_, okByAlias := matchFields[im.Alias]
		_, okByName := matchFields[im.Name]
		if okByAlias || okByName {
			ims = append(ims, &render.Import{
				Alias: im.Alias,
				Name:  im.Name,
				Path:  im.Path,
			})
		}
	}
	return &render.File{
		PackageName: f.PackageName,
		Imports:     ims,
		Struct:      rs,
	}, nil
}

func main() {
	ps, err := property.Parse(prefix, properties)
	if err != nil {
		log.Fatal(err)
	}
	rawFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := entity.NewFile(rawFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := parser.ParseTo(filename, f); err != nil {
		log.Fatal(err)
	}

	rf, err := convertForRendering(f, ps)
	if err != nil {
		log.Fatal(err)
	}

	b, err := render.Render(rf)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
}
