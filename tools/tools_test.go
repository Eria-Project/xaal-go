package tools

import (
	"testing"
)

func TestIsValidAddr(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Good UUID",
			args: args{val: "c91c2708-cb08-11e8-9462-38c98646ac9c"},
			want: true,
		},
		{
			name: "Bad UUID",
			args: args{val: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
			want: false,
		},
		{
			name: "Empty string",
			args: args{val: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidAddr(tt.args.val); got != tt.want {
				t.Errorf("IsValidAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomUUID(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if got := GetRandomUUID(); !IsValidAddr(got) {
			t.Errorf("GetRandomUUID() = %v is not a valid UUID", got)
		}
	})
}

func TestPass2key(t *testing.T) {
	const passphrase = "My passphrase"
	const want = "0ff11736cb905a7ca025da21a694a3ad868f641626a6e74409937279081bde1e"
	t.Run("", func(t *testing.T) {
		if got := Pass2key(passphrase); got != want {
			t.Errorf("Pass2key() = %v, want %v", got, want)
		}
	})
}

func TestIsValidDevType(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Good device type",
			args: args{val: "test.basic"},
			want: true,
		},
		{
			name: "Bad device type",
			args: args{val: "testbasic"},
			want: false,
		},
		{
			name: "Empty string",
			args: args{val: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDevType(tt.args.val); got != tt.want {
				t.Errorf("IsValidDevType() = %v, want %v", got, tt.want)
			}
		})
	}
}
