package common

func ExecuteIntcode(state *[]int) {
	for i := 0; true; {
		opcode := (*state)[i]
		switch opcode {
		case 1:
			i += add(state, i)
		case 2:
			i += mult(state, i)
		case 99:
			return
		}
	}
}

func add(state *[]int, i int) int {
	src1 := (*state)[i+1]
	src2 := (*state)[i+2]
	dest := (*state)[i+3]
	val1 := (*state)[src1]
	val2 := (*state)[src2]
	(*state)[dest] = val1 + val2
	return 4
}

func mult(state *[]int, i int) int {
	src1 := (*state)[i+1]
	src2 := (*state)[i+2]
	dest := (*state)[i+3]
	val1 := (*state)[src1]
	val2 := (*state)[src2]
	(*state)[dest] = val1 * val2
	return 4
}
