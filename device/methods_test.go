package device

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func methodTest(d *Device, args map[string]interface{}) map[string]interface{} {
	return nil
}

func TestDevice_AddMethod(t *testing.T) {
	type fields struct {
		Methods map[string]*Method
	}
	type args struct {
		name string
		f    Function
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *Device
	}{
		{
			name:    "Correct",
			fields:  fields{Methods: make(map[string]*Method)},
			args:    args{name: "TestMethod", f: methodTest},
			wantErr: false,
			want: &Device{
				Methods: map[string]*Method{},
			},
		},
		{
			name:    "Missing name",
			fields:  fields{Methods: make(map[string]*Method)},
			args:    args{name: "", f: methodTest},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				Methods: tt.fields.Methods,
			}
			method, err := d.AddMethod(tt.args.name, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.AddMethod() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.wantErr == false {
				tt.want.Methods = map[string]*Method{"TestMethod": method}
				assert.Equal(t, tt.want, d, "they should be equal")
			}
		})
	}
}

func TestDevice_GetMethods(t *testing.T) {
	methods := map[string]*Method{"TestMethod": &Method{Function: methodTest}}
	t.Run("", func(t *testing.T) {
		d := &Device{
			Methods: methods,
		}
		if got := d.GetMethods(); !reflect.DeepEqual(got, methods) {
			t.Errorf("Device.GetMethods() = %v, want %v", got, methods)
		}
	})
}

func TestDevice_GetMethod(t *testing.T) {
	method := &Method{Function: methodTest}
	type fields struct {
		Methods map[string]*Method
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Method
		wantErr bool
	}{
		{
			name:    "Correct",
			fields:  fields{Methods: map[string]*Method{"TestMethod": method}},
			args:    args{name: "TestMethod"},
			wantErr: false,
			want:    method,
		},
		{
			name:    "Incorrect name",
			fields:  fields{Methods: map[string]*Method{"TestMethod": method}},
			args:    args{name: "Test2"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				Methods: tt.fields.Methods,
			}
			got, err := d.GetMethod(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.GetMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "they should be equal")
		})
	}
}
