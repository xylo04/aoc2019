package common

import (
	"testing"
)

func TestAddAndMult(t *testing.T) {
	testData := map[string]string{
		"1,0,0,0,99":          "2,0,0,0,99",
		"2,3,0,3,99":          "2,3,0,6,99",
		"2,4,4,5,99,0":        "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99": "30,1,1,4,2,5,6,0,99",
	}
	for inputStr, expectedStr := range testData {
		mem, _ := AizuArray(inputStr, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)
		expected, _ := AizuArray(expectedStr, ",")

		NewIntcode(mem, in, out).Execute()
		for i := 0; i < len(mem); i++ {
			if mem[i] != expected[i] {
				t.Errorf("For input %s, final state was expected to be %d but was actually %d",
					inputStr, expected, mem)
				break
			}
		}
	}
}

func TestEcho(t *testing.T) {
	program := "3,0,4,0,99"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != i {
			t.Errorf("Echo intcode test for %d was actually %d", i, actual)
			break
		}
	}
}

func TestParamMode(t *testing.T) {
	mem, _ := AizuArray("1002,4,3,4,33", ",")
	in := make(chan int, 1)
	out := make(chan int, 1)
	expected, _ := AizuArray("1002,4,3,4,99", ",")
	NewIntcode(mem, in, out).Execute()
	for i := 0; i < len(mem); i++ {
		if mem[i] != expected[i] {
			t.Errorf("Parameter mode intcode test failed, should be %d but was %d", expected, mem)
			break
		}
	}
}

func TestEqualTo8_PosMode(t *testing.T) {
	program := "3,9,8,9,10,9,4,9,99,-1,8"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)
		expected := 0
		if i == 8 {
			expected = 1
		}

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != expected {
			t.Errorf("%d", actual)
			break
		}
	}
}

func TestLessThan8_PosMode(t *testing.T) {
	program := "3,9,7,9,10,9,4,9,99,-1,8"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)
		expected := 0
		if i < 8 {
			expected = 1
		}

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != expected {
			t.Errorf("%d", actual)
			break
		}
	}
}
func TestEqualTo8_ImmediateMode(t *testing.T) {
	program := "3,3,1108,-1,8,3,4,3,99"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		expected := 0
		in := make(chan int, 1)
		out := make(chan int, 1)
		if i == 8 {
			expected = 1
		}

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != expected {
			t.Errorf("%d", actual)
			break
		}
	}
}

func TestLessThan8_ImmediateMode(t *testing.T) {
	program := "3,3,1107,-1,8,3,4,3,99"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)
		expected := 0
		if i < 8 {
			expected = 1
		}

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != expected {
			t.Errorf("%d", actual)
			break
		}
	}
}

func TestJump_Pos(t *testing.T) {
	program := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)
		expected := 1
		if i == 0 {
			expected = 0
		}

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != expected {
			t.Errorf("%d", actual)
			break
		}
	}

}

func TestJump_Immediate(t *testing.T) {
	program := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		in := make(chan int, 1)
		out := make(chan int, 1)
		expected := 1
		if i == 0 {
			expected = 0
		}

		in <- i
		go NewIntcode(mem, in, out).Execute()
		actual := <-out

		if actual != expected {
			t.Errorf("%d", actual)
			break
		}
	}

}
