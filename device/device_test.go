package device

import (
	"reflect"
	"testing"
	"xaal-go/tools"
)

var xAALBcastAddr = "00000000-0000-0000-0000-000000000000"
var aliveTimer = uint16(60)

func init() {
	Init(xAALBcastAddr, aliveTimer)
}

func TestNew(t *testing.T) {
	type args struct {
		devType string
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *Device
		wantErr bool
	}{
		{
			name:    "Manual address",
			args:    args{devType: "test.basic", address: "3cd47760-ce4f-11e8-a044-406c8f5172cb"},
			want:    &Device{DevType: "test.basic", Attributes: make(map[string]*Attribute), Address: "3cd47760-ce4f-11e8-a044-406c8f5172cb"},
			wantErr: false,
		},
		{
			name:    "Wrong address",
			args:    args{devType: "test.basic", address: "xxx"},
			wantErr: true,
		},
		{
			name:    "Wrong devType",
			args:    args{devType: "xxx", address: "3cd47760-ce4f-11e8-a044-406c8f5172cb"},
			wantErr: true,
		},
		{
			name:    "Broadcast address",
			args:    args{devType: "test.basic", address: "00000000-0000-0000-0000-000000000000"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.devType, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test auto address
	got, err := New("test.basic", "")
	if err != nil {
		t.Errorf("New() Auto address failed: %v", err)
	}
	if got.Address == "" {
		t.Errorf("New() Auto address no generated")
	}
	if !tools.IsValidAddr(got.Address) {
		t.Errorf("New() Wrong generated address %s", got.Address)
	}

}

func TestDevice_GetTimeout(t *testing.T) {
	type fields struct {
		alivePeriod uint16
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Correct",
			fields: fields{alivePeriod: 60},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				alivePeriod: tt.fields.alivePeriod,
			}
			if got := d.GetTimeout(); got >= tt.fields.alivePeriod {
				t.Errorf("Device.GetTimeout() = %v", got)
			}
		})
	}
}
