package main

import (
	mapset "github.com/deckarep/golang-set"
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

func Test_canSeeEachOther(t *testing.T) {
	type args struct {
		source    point
		dest      point
		asteroids mapset.Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"directly adjacent",
			args{
				point{3, 4},
				point{5, 4},
				mapset.NewSetWith(point{3, 4}, point{5, 4}),
			},
			true,
		},
		{
			"reversed directly adjacent",
			args{
				point{4, 4},
				point{5, 4},
				mapset.NewSetWith(point{3, 4}, point{5, 4}),
			},
			true,
		},
		{
			"adjacent with blocker",
			args{
				point{3, 4},
				point{5, 4},
				mapset.NewSetWith(point{3, 4}, point{4, 4}, point{5, 4}),
			},
			false,
		},
		{
			"reversed adjacent with blocker",
			args{
				point{5, 4},
				point{3, 4},
				mapset.NewSetWith(point{3, 4}, point{4, 4}, point{5, 4}),
			},
			false,
		},
		{
			"directly stacked",
			args{
				point{3, 3},
				point{3, 5},
				mapset.NewSetWith(point{3, 3}, point{3, 5}),
			},
			true,
		},
		{
			"reversed directly stacked",
			args{
				point{3, 5},
				point{3, 3},
				mapset.NewSetWith(point{3, 3}, point{3, 5}),
			},
			true,
		},
		{
			"stacked with blocker",
			args{
				point{3, 3},
				point{3, 5},
				mapset.NewSetWith(point{3, 3}, point{3, 4}, point{3, 5}),
			},
			false,
		},
		{
			"reversed stacked with blocker",
			args{
				point{3, 5},
				point{3, 3},
				mapset.NewSetWith(point{3, 3}, point{3, 4}, point{3, 5}),
			},
			false,
		},
		{
			"down diagonal no blocker",
			args{
				point{2, 2},
				point{4, 4},
				mapset.NewSetWith(point{2, 2}, point{4, 4}),
			},
			true,
		},
		{
			"reversed down diagonal no blocker",
			args{
				point{4, 4},
				point{2, 2},
				mapset.NewSetWith(point{2, 2}, point{4, 4}),
			},
			true,
		},
		{
			"up diagonal no blocker",
			args{
				point{2, 4},
				point{4, 2},
				mapset.NewSetWith(point{2, 4}, point{4, 2}),
			},
			true,
		},
		{
			"reversed up diagonal no blocker",
			args{
				point{4, 2},
				point{2, 4},
				mapset.NewSetWith(point{2, 4}, point{4, 2}),
			},
			true,
		},
		{
			"down diagonal with blocker",
			args{
				point{2, 2},
				point{4, 4},
				mapset.NewSetWith(point{2, 2}, point{3, 3}, point{4, 4}),
			},
			false,
		},
		{
			"reversed down diagonal with blocker",
			args{
				point{4, 4},
				point{2, 2},
				mapset.NewSetWith(point{2, 2}, point{3, 3}, point{4, 4}),
			},
			false,
		},
		{
			"up diagonal with blocker",
			args{
				point{2, 4},
				point{4, 2},
				mapset.NewSetWith(point{2, 4}, point{3, 3}, point{4, 2}),
			},
			false,
		},
		{
			"reversed up diagonal with blocker",
			args{
				point{4, 2},
				point{2, 4},
				mapset.NewSetWith(point{2, 4}, point{3, 3}, point{4, 2}),
			},
			false,
		},
		{
			"45 with blocker",
			args{
				point{0, 0},
				point{2, 4},
				mapset.NewSetWith(point{0, 0}, point{1, 2}, point{2, 4}),
			},
			false,
		}, {
			"reversed 45 with blocker",
			args{
				point{2, 4},
				point{0, 0},
				mapset.NewSetWith(point{0, 0}, point{1, 2}, point{2, 4}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSeeEachOther(tt.args.source, tt.args.dest, tt.args.asteroids); got != tt.want {
				t.Errorf("canSeeEachOther() = %v, want %v", got, tt.want)
			}
		})
	}
}
