package main

import (
	"reflect"
	"testing"
)

func Test_parseFile(t *testing.T) {
	type args struct {
		input string
		wid   int
		hei   int
	}
	tests := []struct {
		name string
		args args
		want spaceImage
	}{
		{
			"case1",
			args{"123456789012", 3, 2},
			spaceImage{
				[][][]int{
					{
						{1, 2, 3},
						{4, 5, 6},
					}, {
						{7, 8, 9},
						{0, 1, 2},
					},
				},
				[]map[int]int{
					{
						1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1,
					},
					{
						7: 1, 8: 1, 9: 1, 0: 1, 1: 1, 2: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFile(tt.args.input, tt.args.wid, tt.args.hei); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checksum(t *testing.T) {
	type args struct {
		image spaceImage
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				spaceImage{
					[][][]int{
						{
							{1, 2, 3},
							{4, 5, 6},
						}, {
							{7, 8, 9},
							{0, 1, 2},
						},
					},
					[]map[int]int{
						{
							1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1,
						},
						{
							7: 1, 8: 1, 9: 1, 0: 1, 1: 1, 2: 1,
						},
					},
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksum(tt.args.image); got != tt.want {
				t.Errorf("checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
