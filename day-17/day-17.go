package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

type VM struct {
	regA         uint64
	regB         uint64
	regC         uint64
	instructions []int
	pc           int
	buffer       string
}

func stoi(s string) int {
	if res, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return res
	}

	return -1
}

func (vm *VM) createCopy() VM {
	instructions := make([]int, len(vm.instructions))
	copy(instructions, vm.instructions)

	return VM{
		regA:         vm.regA,
		regB:         vm.regB,
		regC:         vm.regC,
		instructions: vm.instructions,
		pc:           vm.pc,
		buffer:       vm.buffer,
	}
}

func (vm *VM) getComboValue(operand int) uint64 {
	if operand <= 3 {
		return uint64(operand)
	}
	switch operand {
	case 4:
		return vm.regA
	case 5:
		return vm.regB
	case 6:
		return vm.regC
	case 7:
		log.Fatal("Combo value 7 is preserved!")
	}

	return 0
}

func (vm *VM) exAdv(operand int) {
	numerator := vm.regA
	denumerator := math.Pow(2, float64(vm.getComboValue(operand)))

	vm.regA = uint64(float64(numerator) / denumerator)
}

func (vm *VM) exBxl(operand int) {
	vm.regB = vm.regB ^ uint64(operand)
}

func (vm *VM) exBst(operand int) {
	vm.regB = vm.getComboValue(operand) % 8
}

func (vm *VM) exJnz(operand int) {
	if vm.regA == 0 {
		return
	}
	vm.pc = operand - 2
}

func (vm *VM) exBxc(_ int) {
	vm.regB = vm.regB ^ vm.regC
}

func (vm *VM) exOut(operand int) {
	value := vm.getComboValue(operand) % 8
	if len(vm.buffer) > 0 {
		vm.buffer += ","
	}
	vm.buffer += strconv.FormatInt(int64(value), 10)
}

func (vm *VM) exBdv(operand int) {
	numerator := vm.regA
	denumerator := math.Pow(2, float64(vm.getComboValue(operand)))

	vm.regB = uint64(float64(numerator) / denumerator)
}

func (vm *VM) exCdv(operand int) {
	numerator := vm.regA
	denumerator := math.Pow(2, float64(vm.getComboValue(operand)))

	vm.regC = uint64(float64(numerator) / denumerator)
}

func (vm *VM) tick() bool {
	if vm.pc >= len(vm.instructions) {
		return false
	}

	opcode := vm.instructions[vm.pc]
	operand := vm.instructions[vm.pc+1]

	switch opcode {
	case adv:
		vm.exAdv(operand)
	case bxl:
		vm.exBxl(operand)
	case bst:
		vm.exBst(operand)
	case jnz:
		vm.exJnz(operand)
	case bxc:
		vm.exBxc(operand)
	case out:
		vm.exOut(operand)
	case bdv:
		vm.exBdv(operand)
	case cdv:
		vm.exCdv(operand)
	}

	vm.pc += 2

	return true
}

func partOne(vm VM) {
	vm = vm.createCopy()
	for vm.tick() {
	}

	fmt.Println("Part 1:", vm.buffer)
}

type VMState struct {
	regA int
	regB int
	regC int
	pc   int
}

/*
Search function is specific to my input, we have to reverse engineer the given
instructions to see that at each iterations, the B and C value can be calculated
using the last 3 bits of A: specifically, B can be calculated using those 3
bits, which then can be used to calculate C, which relies only on bits that
are on the left of B, and then A is shifted right by 3 bit, discarding the most recent
B. So, to reconstruct A, we can search for a set of 3 bits at a time, check if
the current A is valid for this new set to generate the wanted instruction

# The given instructions can be translated as follow

2,4, regB = regA % 8 ; get the last 3 bits of A
1,3, regB = regB XOR 3 ; mask with 3

7,5, regC = regA / (2 ^ regB) ; shift A to the right by B bits to get C
; the above line means that C only depends on the bits that are on the left of B

0,3, regA = regA / 8 ; shift A to the right by 3 bits

1,5, regB = regB XOR 5 ; mask B with 5
4,1, regB = regB XOR regC ; mask B with C. Since we know that B and C only relies
; on bits that are on the left of the bits that create B, we can ignore bits that
; are on its right. This is how I come to the conclusion that A can be reconstructed

5,5, OUT regB % 8 ; print the last 3 bits of B => for each instruction, we need to find a set
; of three bits that are left most of A at the position that create that instruction
3,0, QUIT if regA = 0
*/
func search(curr int64, instructions []int, i int) (bool, int64) {
	if i == -1 {
		return true, curr
	}
	ins := instructions[i]
	curr = (curr << 3)
	for regB := range int64(8) {
		tmp := curr ^ regB
		regB ^= 3
		regC := tmp >> regB
		regB ^= 5
		regB ^= regC
		regB &= 7

		if ins == int(regB) {
			if valid, res := search(tmp, instructions, i-1); valid {
				return true, res
			}
		}
	}

	return false, -1
}

func partTwo(vm VM) {
	_, res := search(0, vm.instructions, len(vm.instructions)-1)
	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-17.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pattern, _ := regexp.Compile("[0-9]+")
	vm := VM{pc: 0}
	var match [][]byte

	// Reg A
	scanner.Scan()
	match = pattern.FindSubmatch([]byte(scanner.Text()))
	vm.regA = uint64(stoi(string(match[0])))

	// Reg B
	scanner.Scan()
	match = pattern.FindSubmatch([]byte(scanner.Text()))
	vm.regB = uint64(stoi(string(match[0])))

	// Reg C
	scanner.Scan()
	match = pattern.FindSubmatch([]byte(scanner.Text()))
	vm.regC = uint64(stoi(string(match[0])))

	// Instructions
	scanner.Scan()
	scanner.Scan()
	instructions := pattern.FindAllSubmatch([]byte(scanner.Text()), -1)
	vm.instructions = make([]int, len(instructions))
	for i, ins := range instructions {
		vm.instructions[i] = stoi(string(ins[0]))
	}

	partOne(vm)
	partTwo(vm)
}
