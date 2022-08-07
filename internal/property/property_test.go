package property

import (
	"reflect"
	"testing"
)

func TestParseFieldPropertysProperty(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    []*FieldProperty
		wantErr bool
	}{
		{
			name: "no properties",
			args: args{
				p: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "correct properties for one field (required)",
			args: args{
				p: "SomeFieldProperty:required",
			},
			want: []*FieldProperty{
				{
					Name:     "SomeFieldProperty",
					Required: true,
				},
			},
			wantErr: false,
		},
		{
			name: "correct properties for one field (alias)",
			args: args{
				p: "SomeFieldProperty:alias=AliasForSomeFieldProperty",
			},
			want: []*FieldProperty{
				{
					Name:  "SomeFieldProperty",
					Alias: "AliasForSomeFieldProperty",
				},
			},
			wantErr: false,
		},
		{
			name: "correct properties for more than one fields",
			args: args{
				p: "FirstFieldProperty:required;SomeFieldProperty:alias=SecondFieldProperty",
			},
			want: []*FieldProperty{
				{
					Name:     "FirstFieldProperty",
					Required: true,
				},
				{
					Name:  "SomeFieldProperty",
					Alias: "SecondFieldProperty",
				},
			},
		},
		{
			name: "incorrect property (no property was passed)",
			args: args{
				p: "SomeFieldProperty:",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "incorrect property (no field name and property)",
			args: args{
				p: ":",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "incorrect property (trash)",
			args: args{
				p: "falksdjfalskdjf;aklsdjf;laksjdfasdfasdfjasldf",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "incorrect property (to many ':' symbol)",
			args: args{
				p: "A:A:A:A:A:A:A:A:A:A:A:A;A:",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := ParseFieldsProperty(tt.args.p)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ParseFieldPropertysProperty() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ParseFieldPropertysProperty() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
