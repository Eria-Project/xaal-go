// Package message handle the message struct
package message

import (
	"reflect"
	"testing"
)

var stackVersion = "0.5"

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Message
	}{
		{
			want: Message{Version: "0.5", Targets: []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(stackVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsRequest(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "request"}},
			want:   true,
		},
		{
			name:   "FALSE",
			fields: fields{Header: Header{MsgType: "xxx"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsRequest(); got != tt.want {
				t.Errorf("Message.IsRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsReply(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "reply"}},
			want:   true,
		},
		{
			name:   "FALSE",
			fields: fields{Header: Header{MsgType: "xxx"}},
			want:   false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsReply(); got != tt.want {
				t.Errorf("Message.IsReply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsNotify(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "notify"}},
			want:   true,
		},
		{
			name:   "FALSE",
			fields: fields{Header: Header{MsgType: "xxx"}},
			want:   false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsNotify(); got != tt.want {
				t.Errorf("Message.IsNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsAlive(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "notify", Action: "alive"}},
			want:   true,
		},
		{
			name:   "FALSE type",
			fields: fields{Header: Header{MsgType: "xxx", Action: "alive"}},
			want:   false,
		},
		{
			name:   "FALSE action",
			fields: fields{Header: Header{MsgType: "notify", Action: "xxx"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsAlive(); got != tt.want {
				t.Errorf("Message.IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsAttributesChange(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "notify", Action: "attributesChange"}},
			want:   true,
		},
		{
			name:   "FALSE type",
			fields: fields{Header: Header{MsgType: "xxx", Action: "attributesChange"}},
			want:   false,
		},
		{
			name:   "FALSE action",
			fields: fields{Header: Header{MsgType: "notify", Action: "xxx"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsAttributesChange(); got != tt.want {
				t.Errorf("Message.IsAttributesChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsGetAttributeReply(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "reply", Action: "getAttributes"}},
			want:   true,
		},
		{
			name:   "FALSE type",
			fields: fields{Header: Header{MsgType: "xxx", Action: "getAttributes"}},
			want:   false,
		},
		{
			name:   "FALSE action",
			fields: fields{Header: Header{MsgType: "reply", Action: "xxx"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsGetAttributeReply(); got != tt.want {
				t.Errorf("Message.IsGetAttributeReply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_IsGetDescriptionReply(t *testing.T) {
	type fields struct {
		Header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TRUE",
			fields: fields{Header: Header{MsgType: "reply", Action: "getDescription"}},
			want:   true,
		},
		{
			name:   "FALSE type",
			fields: fields{Header: Header{MsgType: "xxx", Action: "getDescription"}},
			want:   false,
		},
		{
			name:   "FALSE action",
			fields: fields{Header: Header{MsgType: "reply", Action: "xxx"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header: tt.fields.Header,
			}
			if got := m.IsGetDescriptionReply(); got != tt.want {
				t.Errorf("Message.IsGetDescriptionReply() = %v, want %v", got, tt.want)
			}
		})
	}
}
