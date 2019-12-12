package common

import (
	"log"
	"math"
	"sync"
)

type Intcode struct {
	mem     []int
	ptr     int
	in      <-chan int
	out     chan<- int
	wg      *sync.WaitGroup
	relBase int
}

func NewIntcode(mem []int, in <-chan int, out chan<- int) *Intcode {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	return &Intcode{mem, 0, in, out, wg, 0}
}

func NewIntcodeWithWg(mem []int, in <-chan int, out chan<- int, wg *sync.WaitGroup) *Intcode {
	return &Intcode{mem, 0, in, out, wg, 0}
}

func (i *Intcode) Execute() {
	defer i.wg.Done()
programLoop:
	for i.ptr = 0; true; {
		opcode := i.mem[i.ptr] % 100
		switch opcode {
		case 1:
			i.add()
		case 2:
			i.mult()
		case 3:
			i.store(<-i.in)
		case 4:
			i.out <- i.output()
		case 5:
			i.jumpIfTrue()
		case 6:
			i.jumpIfFalse()
		case 7:
			i.lessThan()
		case 8:
			i.equals()
		case 9:
			i.adjustRelBase()
		case 99:
			break programLoop
		default:
			log.Panicf("Unexpected opcode %d", opcode)
		}
	}
	close(i.out)
}

// Read an address from the given parameter
func (i *Intcode) posParam(num int) int {
	mode := i.paramMode(num)
	var pos int
	switch mode {
	case 0:
		// position mode
		pos = i.mem[i.ptr+num]
	case 2:
		pos = i.relBase + i.mem[i.ptr+num]
	default:
		log.Panicf("Unexpected mode %d", mode)
	}
	return pos
}

// Read a value, either referenced by the given parameter's address or directly at the parameter's address,
// depending on the opcode's parameter mode.
func (i *Intcode) valueParam(num int) int {
	mode := i.paramMode(num)
	var pos int
	switch mode {
	case 0:
		// position mode
		pos = i.posParam(num)
	case 1:
		// immediate mode
		pos = i.ptr + num
	case 2:
		pos = i.relBase + i.mem[i.ptr+num]
	default:
		log.Panicf("Unexpected mode %d", mode)
	}
	val := i.mem[pos]
	return val
}

func (i *Intcode) paramMode(num int) interface{} {
	opcode := i.mem[i.ptr]
	mask := int(math.Pow(10, float64(num+1)))
	mode := (opcode / mask) % 10
	return mode
}

func (i *Intcode) add() {
	val1 := i.valueParam(1)
	val2 := i.valueParam(2)
	dest := i.posParam(3)
	i.mem[dest] = val1 + val2
	i.ptr += 4
}

func (i *Intcode) mult() {
	val1 := i.valueParam(1)
	val2 := i.valueParam(2)
	dest := i.posParam(3)
	i.mem[dest] = val1 * val2
	i.ptr += 4
}

func (i *Intcode) store(val int) {
	dest := i.posParam(1)
	i.mem[dest] = val
	i.ptr += 2
}

func (i *Intcode) output() int {
	val := i.valueParam(1)
	i.ptr += 2
	return val
}

func (i *Intcode) jumpIfTrue() {
	compare := i.valueParam(1)
	if compare != 0 {
		i.ptr = i.valueParam(2)
	} else {
		i.ptr += 3
	}
}

func (i *Intcode) jumpIfFalse() {
	compare := i.valueParam(1)
	if compare == 0 {
		i.ptr = i.valueParam(2)
	} else {
		i.ptr += 3
	}
}

func (i *Intcode) lessThan() {
	val1 := i.valueParam(1)
	val2 := i.valueParam(2)
	dest := i.posParam(3)
	if val1 < val2 {
		i.mem[dest] = 1
	} else {
		i.mem[dest] = 0
	}
	i.ptr += 4
}

func (i *Intcode) equals() {
	val1 := i.valueParam(1)
	val2 := i.valueParam(2)
	dest := i.posParam(3)
	if val1 == val2 {
		i.mem[dest] = 1
	} else {
		i.mem[dest] = 0
	}
	i.ptr += 4
}

func (i *Intcode) adjustRelBase() {
	val := i.valueParam(1)
	i.relBase += val
	i.ptr += 2
}
