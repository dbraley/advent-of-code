package main

import (
	"fmt"
	"math"
)

type mem []int

var defaultProgram = mem{
	3, 225,
	1, 225, 6, 6,
	1100, 1, 238, 225,
	104, 0, 1101, 82, 10, 225, 101, 94, 44, 224, 101, -165, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1102, 35, 77, 225, 1102, 28, 71, 225, 1102, 16, 36, 225, 102, 51, 196, 224, 101, -3468, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1001, 48, 21, 224, 101, -57, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 2, 188, 40, 224, 1001, 224, -5390, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1101, 9, 32, 224, 101, -41, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1102, 66, 70, 225, 1002, 191, 28, 224, 101, -868, 224, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 224, 223, 223, 1, 14, 140, 224, 101, -80, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1102, 79, 70, 225, 1101, 31, 65, 225, 1101, 11, 68, 225, 1102, 20, 32, 224, 101, -640, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 329, 101, 1, 223, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 359, 101, 1, 223, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 374, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 389, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 434, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 1001, 223, 1, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 464, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 479, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 509, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 539, 101, 1, 223, 223, 1108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 569, 101, 1, 223, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 599, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 629, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 644, 1001, 223, 1, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

func main() {
	output, err := Run(defaultProgram)
	fmt.Println("Done", err, output)
}

func RunWithInput(noun, verb int, initMem []int) ([]int, error) {
	// Clone the slice
	program := mem(append(initMem[:0:0], initMem...))
	_ = program.put(1, noun)
	_ = program.put(2, verb)
	return Run(program)
}

// run should do a thing
func Run(program mem) ([]int, error) {
	for i := 0; i < len(program); {
		fmt.Printf("Executing instruction at index %v\n", i)
		if shouldContinue, err := ShouldContinue(program, i); !shouldContinue || err != nil {
			return program, err
		}
		p, nextIndex, err := ExecuteInstruction2(program, i)
		if err != nil {
			return program, err
		}
		program = p
		i = nextIndex
	}
	return program, nil
}

func ShouldContinue(program mem, index int) (bool, error) {
	instruction, err := program.get(index)
	if err != nil {
		return false, err
	}
	return instruction != 99, nil
}

func ExecuteInstruction(program mem, index int) ([]int, error) {
	newProgram, _, err := ExecuteInstruction2(program, index)
	return newProgram, err
}

func ExecuteInstruction2(program mem, index int) ([]int, int, error) {
	instruction := program[index]
	opCode := instruction % 100
	modes := instruction / 100
	switch opCode {
	case 1:
		a1, e := getParam(program, modes, index, 0)
		if e != nil {
			return program, -1, e
		}
		a2, e := getParam(program, modes, index, 1)
		if e != nil {
			return program, -1, e
		}

		//fmt.Printf("A(%v): %v + %v = %v\n", index, a1, a2, a1+a2)
		return program, index + 4, putByPtr(program, index+3, a1+a2)
	case 2:
		a1, e := getParam(program, modes, index, 0)
		if e != nil {
			return program, -1, e
		}
		a2, e := getParam(program, modes, index, 1)
		if e != nil {
			return program, -1, e
		}

		//fmt.Printf("M(%v): %v * %v = %v\n", index, a1, a2, a1*a2)
		return program, index + 4, putByPtr(program, index+3, a1*a2)
	case 3:
		input, err := readInput()
		if err != nil {
			return program, -1, err
		}
		return program, index + 2, putByPtr(program, index+1, input)
	case 4:
		output, err := getParam(program, modes, index, 0)
		if err != nil {
			return program, -1, err
		}
		return program, index + 2, writeOutput(output)
	case 5:
		output, err := getParam(program, modes, index, 0)
		if err != nil {
			return program, -1, err
		}
		if output != 0 {
			addr, err := getParam(program, modes, index, 1)
			if err != nil {
				return program, -1, err
			}
			return program, addr, nil
		}
		return program, index + 3, nil
	case 6:
		output, err := getParam(program, modes, index, 0)
		if err != nil {
			return program, -1, err
		}
		if output == 0 {
			addr, err := getParam(program, modes, index, 1)
			if err != nil {
				return program, -1, err
			}
			return program, addr, nil
		}
		return program, index + 3, nil
	case 7:
		p1, err := getParam(program, modes, index, 0)
		if err != nil {
			return program, -1, err
		}
		p2, err := getParam(program, modes, index, 1)
		if err != nil {
			return program, -1, err
		}
		out := 0
		if p1 < p2 {
			out = 1
		}
		return program, index + 4, putByPtr(program, index+3, out)
	case 8:
		p1, err := getParam(program, modes, index, 0)
		if err != nil {
			return program, -1, err
		}
		p2, err := getParam(program, modes, index, 1)
		if err != nil {
			return program, -1, err
		}
		out := 0
		if p1 == p2 {
			out = 1
		}
		return program, index + 4, putByPtr(program, index+3, out)
	default:
		return program, -1, fmt.Errorf("illegal opCode %v at index %v\n", opCode, index)
	}
}

// TODO: This should be passed in
func readInput() (int, error) {
	//return 1, nil
	return 5, nil
}

func writeOutput(output int) error {
	_, err := fmt.Printf("output: %d\n", output)
	return err
}

func getParam(program mem, modes int, index int, paramIndex int) (int, error) {
	if isPositionMode(paramIndex, modes) {
		//fmt.Printf("isPositionaMode = true, calling getByPtr(%v)\n", index + paramIndex+1)
		return getByPtr(program, index+paramIndex+1)
	}
	//fmt.Printf("isPositionaMode = false, calling get(%v)\n", index + paramIndex+1)
	return program.get(index + paramIndex + 1)
}

func isPositionMode(paramIndex, modes int) bool {
	//fmt.Printf("checking param index %d against modes %d: %d\n", paramIndex, modes, (modes/(10 ^ paramIndex)) % 10)
	return ((modes / int(math.Pow10(paramIndex))) % 10) == 0
}

func getByPtr(program mem, index int) (int, error) {
	ptr, err := program.get(index)
	if err != nil {
		return 0, err
	}
	val, err := program.get(ptr)
	if err != nil {
		return 0, err
	}
	//fmt.Printf("Fetching index %v -> %v -> %v\n", index, ptr, val)
	return val, nil
}

func putByPtr(program mem, index, value int) error {
	ptr, err := program.get(index)
	if err != nil {
		return err
	}
	//oldVal, _ := program.get(ptr)
	//fmt.Printf("Modifying index %v->%v from %v to %v\n", index, ptr, oldVal, value)
	return program.put(ptr, value)
}

func (a mem) get(i int) (int, error) {
	if i >= 0 && i < len(a) {
		return a[i], nil
	}
	return -1, fmt.Errorf("index %v out of bounds", i)
}

func (a mem) put(i, val int) error {
	if i >= 0 && i < len(a) {
		a[i] = val
		return nil
	}
	return fmt.Errorf("index %v out of bounds", i)
}
