package device

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDevice_SetAttributeValue(t *testing.T) {
	type fields struct {
		Attributes map[string]*Attribute
	}
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{Attributes: map[string]*Attribute{"test": &Attribute{Name: "test", defaultValue: 0}}},
			args:   args{name: "test", value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, _ := New("test.basic", "")
			d.Attributes = tt.fields.Attributes
			d.SetAttributeValue(tt.args.name, tt.args.value)
			if d.Attributes["test"].Value != tt.args.value {
				t.Errorf("TestDevice_SetAttributeValue The attribute hasn't been set")
			}
		})
	}
}

func TestDevice_AddUnsupportedAttribute(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, _ := New("test.basic", "")
			d.AddUnsupportedAttribute(tt.args.name)
			if len(d.unsupportedAttributes) == 0 {
				t.Errorf("TestDevice_AddUnsupportedAttribute The attribute hasn't been added to the list")
			}
			if d.unsupportedAttributes[0] != tt.args.name {
				t.Errorf("TestDevice_AddUnsupportedAttribute The attribute doesn't seems to match")
			}
		})
	}
}

func TestDevice_NewAttribute(t *testing.T) {
	type args struct {
		name         string
		defaultValue interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Attribute
		wantErr bool
	}{
		{
			name:    "Correct name",
			args:    args{name: "Test", defaultValue: 0},
			wantErr: false,
		},
		{
			name:    "Empty name",
			args:    args{name: "", defaultValue: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, _ := New("test.basic", "")
			got, err := d.NewAttribute(tt.args.name, tt.args.defaultValue)
			fmt.Println(err != nil)
			fmt.Println((err != nil) != tt.wantErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				tt.want = &Attribute{Name: "Test", defaultValue: 0, Device: d}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
