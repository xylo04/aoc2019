package common

func ExecuteIntcode(mem *[]int, input int) int {
	out := 0
programLoop:
	for i := 0; true; {
		opcode := (*mem)[i]
		switch opcode {
		case 1:
			add(mem, &i)
		case 2:
			mult(mem, &i)
		case 3:
			store(input, mem, &i)
		case 4:
			out = output(mem, &i)
		case 99:
			break programLoop
		}
	}
	return out
}

func add(mem *[]int, i *int) {
	src1 := (*mem)[*i+1]
	src2 := (*mem)[*i+2]
	dest := (*mem)[*i+3]
	val1 := (*mem)[src1]
	val2 := (*mem)[src2]
	(*mem)[dest] = val1 + val2
	*i += 4
}

func mult(mem *[]int, i *int) {
	src1 := (*mem)[*i+1]
	src2 := (*mem)[*i+2]
	dest := (*mem)[*i+3]
	val1 := (*mem)[src1]
	val2 := (*mem)[src2]
	(*mem)[dest] = val1 * val2
	*i += 4
}

func store(input int, mem *[]int, i *int) {
	panic("Not implemented")
}

func output(mem *[]int, i *int) int {
	panic("Not implemented")
}
