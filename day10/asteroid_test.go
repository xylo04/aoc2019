package main

import (
	"reflect"
	"testing"
)

func TestFindBestAsteroid(t *testing.T) {
	type args struct {
		lines []string
	}
	type returns struct {
		m int
		p point
	}
	tests := []struct {
		name string
		args args
		want returns
	}{
		{
			"case1",
			args{
				lines: []string{
					".#..#",
					".....",
					"#####",
					"....#",
					"...##",
				},
			},
			returns{8, point{3, 4}},
		},
		{
			"case2",
			args{
				lines: []string{
					"......#.#.",
					"#..#.#....",
					"..#######.",
					".#.#.###..",
					".#..#.....",
					"..#....#.#",
					"#..#....#.",
					".##.#..###",
					"##...#..#.",
					".#....####",
				},
			},
			returns{33, point{5, 8}},
		},
		{
			"case3",
			args{
				lines: []string{
					"#.#...#.#.",
					".###....#.",
					".#....#...",
					"##.#.#.#.#",
					"....#.#.#.",
					".##..###.#",
					"..#...##..",
					"..##....##",
					"......#...",
					".####.###.",
				},
			},
			returns{35, point{1, 2}},
		}, {
			"case4",
			args{
				lines: []string{
					".#..#..###",
					"####.###.#",
					"....###.#.",
					"..###.##.#",
					"##.##.#.#.",
					"....###..#",
					"..#.#..#.#",
					"#..#.#.###",
					".##...##.#",
					".....#.#..",
				},
			},
			returns{41, point{6, 3}},
		}, {
			"case5",
			args{
				lines: []string{
					".#..##.###...#######",
					"##.############..##.",
					".#.######.########.#",
					".###.#######.####.#.",
					"#####.##.#.##.###.##",
					"..#####..#.#########",
					"####################",
					"#.####....###.#.#.##",
					"##.#################",
					"#####.##.###..####..",
					"..######..##.#######",
					"####.##.####...##..#",
					".#####..#.######.###",
					"##...#.##########...",
					"#.##########.#######",
					".####.#.###.###.#.##",
					"....##.##.###..#####",
					".#.#.###########.###",
					"#.#.#.#####.####.###",
					"###.##.####.##.#..##",
				},
			},
			returns{210, point{11, 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := parseAsteroids(tt.args.lines)
			gotM, gotP := findBestAsteroid(field)
			if !reflect.DeepEqual(gotP, tt.want.p) || !reflect.DeepEqual(gotM, tt.want.m) {
				t.Errorf("findBestAsteroid() max = %v, want %v", gotM, tt.want.m)
				t.Errorf("findBestAsteroid() point = %v, want %v", gotP, tt.want.p)
			}
		})
	}
}
