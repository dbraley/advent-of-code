package main

import "fmt"

var defaultProgram = mem{
	1, 0, 0, 3,
	1, 1, 2, 3,
	1, 3, 4, 3,
	1, 5, 0, 3,
	2, 1, 10, 19,
	1, 19, 5, 23,
	2, 23, 9, 27,
	1, 5, 27, 31,
	1, 9, 31, 35,
	1, 35, 10, 39,
	2, 13, 39, 43,
	1, 43, 9, 47,
	1, 47, 9, 51,
	1, 6, 51, 55,
	1, 13, 55, 59,
	1, 59, 13, 63,
	1, 13, 63, 67,
	1, 6, 67, 71,
	1, 71, 13, 75,
	2, 10, 75, 79,
	1, 13, 79, 83,
	1, 83, 10, 87,
	2, 9, 87, 91,
	1, 6, 91, 95,
	1, 9, 95, 99,
	2, 99, 10, 103,
	1, 103, 5, 107,
	2, 6, 107, 111,
	1, 111, 6, 115,
	1, 9, 115, 119,
	1, 9, 119, 123,
	2, 10, 123, 127,
	1, 127, 5, 131,
	2, 6, 131, 135,
	1, 135, 5, 139,
	1, 9, 139, 143,
	2, 143, 13, 147,
	1, 9, 147, 151,
	1, 151, 2, 155,
	1, 9, 155, 0,
	99, 2, 0, 14, 0}

type mem []int

func main() {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			out, err := RunWithInput(noun, verb, defaultProgram)
			if err != nil {
				fmt.Printf("ERROR: %02d%02d !!! %v\n", noun, verb, err)
				continue
			}
			if out[0] == 19690720 {
				fmt.Printf("%02d%02d : %d\n", noun, verb, out[0])
			} else if noun == 12 && verb == 2 {
				fmt.Printf("%02d%02d : %d\n", noun, verb, out[0])
			} else {
				//fmt.Printf("%02d%02d : %d\n", noun, verb, out[0])
			}
		}
	}
}

func RunWithInput(noun, verb int, initMem []int) ([]int, error) {
	// Clone the slice
	program := mem(append(initMem[:0:0], initMem...))
	_ = program.put(1, noun)
	_ = program.put(2, verb)
	return Run(program)
}

func Run(program mem) ([]int, error) {
	for i := 0; i < len(program); i += 4 {
		if shouldContinue, err := ShouldContinue(program, i); !shouldContinue || err != nil {
			return program, err
		}
		p, err := ExecuteInstruction(program, i)
		if err != nil {
			return program, err
		}
		program = p
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
	instruction := program[index]
	switch instruction {
	case 1:
		a1, e := getByPtr(program, index+1)
		if e != nil {
			return program, e
		}
		a2, e := getByPtr(program, index+2)
		if e != nil {
			return program, e
		}

		//fmt.Printf("A(%v): %v + %v = %v\n", index, a1, a2, a1+a2)
		return program, putByPtr(program, index+3, a1+a2)
	case 2:
		a1, e := getByPtr(program, index+1)
		if e != nil {
			return program, e
		}
		a2, e := getByPtr(program, index+2)
		if e != nil {
			return program, e
		}

		//fmt.Printf("M(%v): %v * %v = %v\n", index, a1, a2, a1*a2)
		return program, putByPtr(program, index+3, a1*a2)
	default:
		return program, fmt.Errorf("illegal instruction %v\n", instruction)
	}
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
