package main

import "testing"

func Test_personalColor(t *testing.T) {
	type args struct {
		personalData HumanStatus
	}
	tests := []struct {
		name                  string
		args                  args
		wantPersonalColorData string
	}{
		{"test1", args{HumanStatus{"haruki", 12, 31}}, "153,255,255"},
		{"test2", args{HumanStatus{"erina", 1, 1}}, "127,21,8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPersonalColorData := personalColor(tt.args.personalData); gotPersonalColorData != tt.wantPersonalColorData {
				t.Errorf("personalColor() = %v, want %v", gotPersonalColorData, tt.wantPersonalColorData)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
