package common

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
		case 99:
			break programLoop
		}
	}
	return io
}

// Read an address from the given parameter
func posParam(num int) int {
	return (*mem)[ptr+num]
}

// Read a value referenced by the given parameter's address
func valueRefParam(num int) int {
	pos := posParam(num)
	val := (*mem)[pos]
	return val
}

func add() {
	val1 := valueRefParam(1)
	val2 := valueRefParam(2)
	dest := posParam(3)
	(*mem)[dest] = val1 + val2
	ptr += 4
}

func mult() {
	val1 := valueRefParam(1)
	val2 := valueRefParam(2)
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
	val := valueRefParam(1)
	ptr += 2
	return val
}
