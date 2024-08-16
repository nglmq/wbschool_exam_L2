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
		// TODO: Add test cases.
		
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
