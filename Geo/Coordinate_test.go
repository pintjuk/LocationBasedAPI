package Geo

import (
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		name      string
		args      string
		wantError bool
		want1     Coordinate
	}{
		{
			name:      "default",
			args:      "",
			wantError: true,
			want1:     Coordinate{},
		},
		{
			name:      "should parse correct",
			args:      " (0.0,   0.0) ",
			wantError: false,
			want1:     Coordinate{0.0, 0.0},
		},
		{
			name:      "should return error if another symbol then comma is used as devider",
			args:      " (0.0;   0.0) ",
			wantError: true,
			want1:     Coordinate{},
		},
		{
			name:      "should return error if brases are not used",
			args:      " 0.0,   0.0",
			wantError: true,
			want1:     Coordinate{},
		},
		{
			name:      "should parse correct",
			args:      " (10,   10.10) ",
			wantError: false,
			want1:     Coordinate{10, 10.10},
		},
		{
			name:      "should parse toString output",
			args:      Coordinate{10.1234, 3.21}.ToString(),
			wantError: false,
			want1:     Coordinate{10.1234, 3.21},
		},
		{
			name:      "should return error when comma is used as decimpal sign ",
			args:      " (10,   10,10) ",
			wantError: true,
			want1:     Coordinate{},
		},
		{
			name:      "should return error when not floats  ",
			args:      " (10,   aoeu) ",
			wantError: true,
			want1:     Coordinate{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseString(tt.args)
			if (got == nil) == tt.wantError {
				t.Errorf("ParseString() got = %v, want error: %v", got, tt.wantError)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
