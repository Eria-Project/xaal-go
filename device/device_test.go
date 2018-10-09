package device

import (
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
