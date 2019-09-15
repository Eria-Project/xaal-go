package device

import (
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
func TestDevice_getAttributes(t *testing.T) {
	type args struct {
		d    *Device
		args map[string]interface{}
	}
	device := Device{
		Attributes: map[string]*Attribute{
			"a": &Attribute{
				Name:  "a",
				Value: 1,
			},
			"b": &Attribute{
				Name:  "b",
				Value: "y",
			},
		},
	}

	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Nil args",
			args: args{d: &device, args: nil},
			want: map[string]interface{}{
				"a": 1,
				"b": "y",
			},
		},
		{
			name: "Empty args",
			args: args{d: &device, args: make(map[string]interface{})},
			want: map[string]interface{}{
				"a": 1,
				"b": "y",
			},
		},
		{
			name: "Empty 'attributes' array arg",
			args: args{d: &device, args: map[string]interface{}{
				"attributes": []string{},
			}},
			want: map[string]interface{}{
				"a": 1,
				"b": "y",
			},
		},
		{
			name: "Filter attribute 'a' only",
			args: args{d: &device, args: map[string]interface{}{
				"attributes": []string{"a"},
			}},
			want: map[string]interface{}{
				"a": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAttributes(tt.args.d, tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
