package main

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantR   string
		wantErr bool
	}{
		{
			name:    "Test 1",
			args:    args{s: "a4bc2d5e"},
			wantR:   "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "Test 2",
			args:    args{s: "abcd"},
			wantR:   "abcd",
			wantErr: false,
		},
		{
			name:    "Test 3",
			args:    args{s: "45"},
			wantR:   "",
			wantErr: true,
		},
		{
			name:    "Test 4",
			args:    args{s: ""},
			wantR:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := Unpack(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("Unpack() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
