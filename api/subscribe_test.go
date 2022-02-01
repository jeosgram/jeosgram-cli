package api

import (
	"testing"
)

func Test_makeURL(t *testing.T) {
	type args struct {
		deviceID  string
		eventName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"all-all",
			args{
				"",
				"",
			},
			"/v1/devices/events",
		},
		{
			"all-prefix",
			args{
				"",
				"gnss",
			},
			"/v1/devices/events/gnss",
		},
		{
			"device-all",
			args{
				"abc123",
				"",
			},
			"/v1/devices/abc123/events",
		},
		{
			"device-prefix",
			args{
				"abc123",
				"gnss",
			},
			"/v1/devices/abc123/events/gnss",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeURL(tt.args.deviceID, tt.args.eventName); got != tt.want {
				t.Errorf("makeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
