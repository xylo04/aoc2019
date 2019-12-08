package common

import (
	"log"
	"math"
)

var mem *[]int
var ptr int

func ExecuteIntcode(memory *[]int, input int) int {
	mem = memory
	io := input
programLoop:
	for ptr = 0; true; {
		opcode := (*mem)[ptr] % 100
		switch opcode {
		case 1:
			add()
		case 2:
			mult()
		case 3:
			store(io)
		case 4:
			io = output()
		case 5:
			jumpIfTrue()
		case 6:
			jumpIfFalse()
		case 7:
			lessThan()
		case 8:
			equals()
		case 99:
			break programLoop
		default:
			log.Panicf("Unexpected opcode %d", opcode)
		}
	}
	return io
}

// Read an address from the given parameter
func posParam(num int) int {
	return (*mem)[ptr+num]
}

// Read a value, either referenced by the given parameter's address or directly at the parameter's address,
// depending on the opcode's parameter mode.
func valueParam(num int) int {
	mode := paramMode(num)
	var pos int
	switch mode {
	case 0:
		// position mode
		pos = posParam(num)
	case 1:
		// immediate mode
		pos = ptr + num
	default:
		log.Panicf("Unexpected mode %d", mode)
	}
	val := (*mem)[pos]
	return val
}

func paramMode(num int) interface{} {
	opcode := (*mem)[ptr]
	mask := int(math.Pow(10, float64(num+1)))
	mode := (opcode / mask) % 10
	return mode
}

func add() {
	val1 := valueParam(1)
	val2 := valueParam(2)
	dest := posParam(3)
	(*mem)[dest] = val1 + val2
	ptr += 4
}

func mult() {
	val1 := valueParam(1)
	val2 := valueParam(2)
	dest := posParam(3)
	(*mem)[dest] = val1 * val2
	ptr += 4
}

func store(val int) {
	dest := posParam(1)
	(*mem)[dest] = val
	ptr += 2
}

func output() int {
	val := valueParam(1)
	ptr += 2
	return val
}

func jumpIfTrue() {
	compare := valueParam(1)
	if compare != 0 {
		ptr = valueParam(2)
	} else {
		ptr += 3
	}
}

func jumpIfFalse() {
	compare := valueParam(1)
	if compare == 0 {
		ptr = valueParam(2)
	} else {
		ptr += 3
	}
}

func lessThan() {
	val1 := valueParam(1)
	val2 := valueParam(2)
	dest := posParam(3)
	if val1 < val2 {
		(*mem)[dest] = 1
	} else {
		(*mem)[dest] = 0
	}
	ptr += 4
}

func equals() {
	val1 := valueParam(1)
	val2 := valueParam(2)
	dest := posParam(3)
	if val1 == val2 {
		(*mem)[dest] = 1
	} else {
		(*mem)[dest] = 0
	}
	ptr += 4
}
