// Package xaal : Devices Management
package xaal

import (
	"testing"
	"github.com/project-eria/xaal-go/device"
	"github.com/project-eria/xaal-go/message"

	"github.com/stretchr/testify/assert"
)

func Test_filterMsgForDevices(t *testing.T) {
	devices := map[string]*device.Device{
		"1051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
			DevType: "test1.basic1",
			Address: "1051dd80-d93a-11e8-8829-38c98646ac9c",
		},
		"2051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
			DevType: "test2.basic1",
			Address: "2051dd80-d93a-11e8-8829-38c98646ac9c",
		},
		"3051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
			DevType: "test1.basic2",
			Address: "3051dd80-d93a-11e8-8829-38c98646ac9c",
		},
	}
	type args struct {
		msg     *message.Message
		devices map[string]*device.Device
	}
	tests := []struct {
		name string
		args args
		want map[string]*device.Device
	}{
		{
			name: "isAlive, with no devTypes",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "isAlive",
					},
				},
				devices: devices,
			},
			want: nil,
		},
		{
			name: "isAlive, with devTypes = any.any",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "isAlive",
					},
					Body: map[string]interface{}{
						"devTypes": "any.any",
					},
				},
				devices: devices,
			},
			want: devices,
		},
		{
			name: "isAlive, with devTypes = [any.any]",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "isAlive",
					},
					Body: map[string]interface{}{
						"devTypes": []string{"any.any"},
					},
				},
				devices: devices,
			},
			want: devices,
		},
		{
			name: "isAlive, with devTypes = [test1.basic1, test2.basic1]",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "isAlive",
					},
					Body: map[string]interface{}{
						"devTypes": []string{"test1.basic1", "test2.basic1"},
					},
				},
				devices: devices,
			},
			want: map[string]*device.Device{
				"1051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
					DevType: "test1.basic1",
					Address: "1051dd80-d93a-11e8-8829-38c98646ac9c",
				},
				"2051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
					DevType: "test2.basic1",
					Address: "2051dd80-d93a-11e8-8829-38c98646ac9c",
				},
			},
		},
		{
			name: "isAlive, with devTypes = [test1.any]",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "isAlive",
					},
					Body: map[string]interface{}{
						"devTypes": []string{"test1.any"},
					},
				},
				devices: devices,
			},
			want: map[string]*device.Device{
				"1051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
					DevType: "test1.basic1",
					Address: "1051dd80-d93a-11e8-8829-38c98646ac9c",
				},
				"3051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
					DevType: "test1.basic2",
					Address: "3051dd80-d93a-11e8-8829-38c98646ac9c",
				},
			},
		},
		{
			name: "Action, with empty targets (all/broadcast)",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "test",
					},
					Targets: []string{},
				},
				devices: devices,
			},
			want: devices,
		},
		{
			name: "Action, with 1 target (specific)",
			args: args{
				msg: &message.Message{
					Header: message.Header{
						Action: "test",
					},
					Targets: []string{"2051dd80-d93a-11e8-8829-38c98646ac9c"},
				},
				devices: devices,
			},
			want: map[string]*device.Device{
				"2051dd80-d93a-11e8-8829-38c98646ac9c": &device.Device{
					DevType: "test2.basic1",
					Address: "2051dd80-d93a-11e8-8829-38c98646ac9c",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterMsgForDevices(tt.args.msg, tt.args.devices)
			assert.Equal(t, tt.want, got, "they should be equal")
		})
	}
}
