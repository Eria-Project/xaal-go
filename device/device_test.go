package device

import (
	"reflect"
	"testing"
)

func TestDevice_SetDevType(t *testing.T) {
	type fields struct {
		DevType     string
		Address     string
		alivePeriod int
	}
	type args struct {
		devType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Simple device type",
			fields:  fields{},
			args:    args{devType: "test.basic"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Device{
				DevType:     tt.fields.DevType,
				Address:     tt.fields.Address,
				alivePeriod: tt.fields.alivePeriod,
			}
			var err error
			if err = d.SetDevType(tt.args.devType); (err != nil) != tt.wantErr {
				t.Errorf("Device.SetDevType() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if d.DevType != tt.args.devType {
					t.Errorf("TestDevice_SetDevType() = '%v', want '%v'", d.DevType, tt.args.devType)
				}
			}
		})
	}
}

func TestDevice_SetAddress(t *testing.T) {
	type fields struct {
		DevType     string
		Address     string
		alivePeriod int
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Simple device address",
			fields:  fields{},
			args:    args{address: "c91c2708-cb08-11e8-9462-38c98646ac9c"},
			wantErr: false,
		},
		{
			name:    "Broadcast address",
			fields:  fields{},
			args:    args{address: "00000000-0000-0000-0000-000000000000"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Device{
				DevType:     tt.fields.DevType,
				Address:     tt.fields.Address,
				alivePeriod: tt.fields.alivePeriod,
			}
			var err error
			if err = d.SetAddress(tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("Device.SetAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if d.Address != tt.args.address {
					t.Errorf("TestDevice_SetAddress() = '%v', want '%v'", d.Address, tt.args.address)
				}
			}
		})
	}
}

func TestDevice_GetTimeout(t *testing.T) {
	type fields struct {
		DevType     string
		Address     string
		alivePeriod int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "",
			fields: fields{alivePeriod: 60},
			want:   120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Device{
				DevType:     tt.fields.DevType,
				Address:     tt.fields.Address,
				alivePeriod: tt.fields.alivePeriod,
			}
			if got := d.GetTimeout(); got != tt.want {
				t.Errorf("Device.GetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_AddAttribute(t *testing.T) {
	type fields struct {
		DevType               string
		Address               string
		alivePeriod           int
		unsupportedAttributes []string
		attributes            []Attribute
	}
	type args struct {
		attr *Attribute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Integer attribute",
			fields: fields{},
			args:   args{attr: &Attribute{"int", 10, 0}},
		},
		{
			name:   "String attribute",
			fields: fields{},
			args:   args{attr: &Attribute{"string", "A", "B"}},
		},
		{
			name:   "Bool attribute",
			fields: fields{},
			args:   args{attr: &Attribute{"bool", true, false}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				DevType:               tt.fields.DevType,
				Address:               tt.fields.Address,
				alivePeriod:           tt.fields.alivePeriod,
				unsupportedAttributes: tt.fields.unsupportedAttributes,
				attributes:            tt.fields.attributes,
			}
			d.AddAttribute(tt.args.attr)
			if len(d.attributes) == 0 {
				t.Errorf("TestDevice_AddAttribute The attribute hasn't been added to the list")
			}
			if !reflect.DeepEqual(d.attributes[0], *tt.args.attr) {
				t.Errorf("TestDevice_AddAttribute The attribute doesn't seems to match")
			}
		})
	}
}

func TestDevice_NewAttribute(t *testing.T) {
	type fields struct {
		DevType               string
		Address               string
		alivePeriod           int
		unsupportedAttributes []string
		attributes            []Attribute
	}
	type args struct {
		name         string
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Attribute
	}{
		{
			fields: fields{},
			args:   args{name: "Test", defaultValue: 0},
			want:   Attribute{name: "Test", defaultValue: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				DevType:               tt.fields.DevType,
				Address:               tt.fields.Address,
				alivePeriod:           tt.fields.alivePeriod,
				unsupportedAttributes: tt.fields.unsupportedAttributes,
				attributes:            tt.fields.attributes,
			}
			if got := d.NewAttribute(tt.args.name, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Device.NewAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_findAttributeIndex(t *testing.T) {
	type fields struct {
		DevType               string
		Address               string
		alivePeriod           int
		unsupportedAttributes []string
		attributes            []Attribute
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "Existing attribute",
			fields: fields{attributes: []Attribute{Attribute{"test", 1, 0}}},
			args:   args{name: "test"},
			want:   0,
		},
		{
			name:   "Not existing attribute",
			fields: fields{},
			args:   args{name: "test"},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				DevType:               tt.fields.DevType,
				Address:               tt.fields.Address,
				alivePeriod:           tt.fields.alivePeriod,
				unsupportedAttributes: tt.fields.unsupportedAttributes,
				attributes:            tt.fields.attributes,
			}
			if got := d.findAttributeIndex(tt.args.name); got != tt.want {
				t.Errorf("Device.findAttributeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_SetAttributeValue(t *testing.T) {
	type fields struct {
		DevType               string
		Address               string
		alivePeriod           int
		unsupportedAttributes []string
		attributes            []Attribute
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
			fields: fields{attributes: []Attribute{Attribute{"test", nil, 0}}},
			args:   args{name: "test", value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				DevType:               tt.fields.DevType,
				Address:               tt.fields.Address,
				alivePeriod:           tt.fields.alivePeriod,
				unsupportedAttributes: tt.fields.unsupportedAttributes,
				attributes:            tt.fields.attributes,
			}
			d.SetAttributeValue(tt.args.name, tt.args.value)
			if d.attributes[0].value != tt.args.value {
				t.Errorf("TestDevice_SetAttributeValue The attribute hasn't been set")
			}
		})
	}
}

func TestDevice_DelAttribute(t *testing.T) {
	type fields struct {
		DevType               string
		Address               string
		alivePeriod           int
		unsupportedAttributes []string
		attributes            []Attribute
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Empty name",
			fields:  fields{attributes: []Attribute{Attribute{"test", nil, 0}}},
			args:    args{name: ""},
			wantErr: true,
		},
		{
			name:    "Not existing name",
			fields:  fields{attributes: []Attribute{Attribute{"test", nil, 0}}},
			args:    args{name: "a"},
			wantErr: true,
		},
		{
			name:    "Existing name",
			fields:  fields{attributes: []Attribute{Attribute{"test", nil, 0}}},
			args:    args{name: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				DevType:               tt.fields.DevType,
				Address:               tt.fields.Address,
				alivePeriod:           tt.fields.alivePeriod,
				unsupportedAttributes: tt.fields.unsupportedAttributes,
				attributes:            tt.fields.attributes,
			}
			err := d.DelAttribute(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.DelAttribute() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && len(d.attributes) != 0 {
				t.Errorf("Device.DelAttribute() the attribute hasn't been removed from the list")
			}
			if err != nil && len(d.attributes) != 1 {
				t.Errorf("Device.DelAttribute() the attribute has been unexpected removed from the list")
			}
		})
	}
}

func TestDevice_AddUnsupportedAttribute(t *testing.T) {
	type fields struct {
		DevType               string
		Address               string
		alivePeriod           int
		unsupportedAttributes []string
		attributes            []Attribute
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{},
			args:   args{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				DevType:               tt.fields.DevType,
				Address:               tt.fields.Address,
				alivePeriod:           tt.fields.alivePeriod,
				unsupportedAttributes: tt.fields.unsupportedAttributes,
				attributes:            tt.fields.attributes,
			}
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
