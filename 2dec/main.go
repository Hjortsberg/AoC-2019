package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// to run this, you need to paste your comma separated input and hit ctrl + d
// twice to send end of line
func main() {

	// split on comma scanner is example code from go docs for bufio.
	// please note that it's not my work.
	scanner := bufio.NewScanner(os.Stdin)
	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.

	var intCode []int
	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())
		intCode = append(intCode, intVal)
		if err != nil {
			fmt.Println("some strconv error")
		}
	}

nounLoop:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			fmt.Println(verb)
			currIntCode := make([]int, len(intCode))
			copy(currIntCode, intCode)
			currIntCode[1] = noun
			currIntCode[2] = verb
			fmt.Printf("noun: %d\nverb: %d\n", currIntCode[1], currIntCode[2])
			result := intStateMachine(currIntCode)
			if result[0] == 19690720 {
				fmt.Printf("noun: %d\nverb: %d\n", noun, verb)
				break nounLoop
			}
		}
	}
	//fmt.Println(intStateMachine(intCode))
	fmt.Println(intCode[0])
}

func intStateMachine(intCode []int) []int {
	pc := 0
	for pc < len(intCode) {
		opCode := intCode[pc]
		fmt.Printf("opcode: %d\n", opCode)
		fmt.Printf("pc: %d\n", pc)
		if opCode == 99 {
			break
		} else if opCode == 1 {
			newValue := intCode[intCode[pc+1]] + intCode[intCode[pc+2]] // perform numerical add op
			address := intCode[pc+3]                                    // find address to write value to
			intCode[address] = newValue                                 //assign new value to the address found in previous step
			pc += 4
		} else if opCode == 2 {
			newValue := intCode[intCode[pc+1]] * intCode[intCode[pc+2]] // perform numerical mult op
			address := intCode[pc+3]                                    // find address to write value to
			intCode[address] = newValue                                 //assign new value to the address found in previous step
			pc += 4
		}
	}
	return intCode
}
