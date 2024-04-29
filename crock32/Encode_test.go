package crock32

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want string
	}{
		{
			"Integer: 1",
			1,
			"04",
		},
		{
			"Integer: 0001",
			0001,
			"04",
		},
		{
			"Integer: 1111",
			1111,
			"0HBG",
		},
		{
			"Integer: 123456",
			123456,
			"07H40",
		},
		{
			"Timestamp: Mon Apr 29 06:42:35 2024 UTC",
			1714372955,
			"CRQM2PR",
		},
		{
			"Integer: 1234567890",
			1234567890,
			"96B05MG",
		},
		{
			"Integer: 1234567890123456789",
			1234567890123456789,
			"24H11X3XX60HA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.arg); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
			if got := Decode(tt.want); got != tt.arg {
				t.Errorf("Decode() = %v, want %v", got, tt.arg)
			}
		})
	}
}
