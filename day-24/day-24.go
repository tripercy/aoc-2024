package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Gate func(int, int) int

type Wire struct {
	name     string
	inputA   *Wire
	inputB   *Wire
	gate     Gate
	val      int
	gateName string
}

var gateStr = map[string]Gate{
	"AND": func(i1, i2 int) int { return i1 & i2 },
	"OR":  func(i1, i2 int) int { return i1 | i2 },
	"XOR": func(i1, i2 int) int { return i1 ^ i2 },
}

var wires = make(map[string]*Wire)

func getWire(name string) *Wire {
	if wire, exists := wires[name]; !exists {
		wires[name] = &Wire{
			name:   name,
			val:    -1,
			inputA: nil,
			inputB: nil,
			gate:   nil,
		}
		return wires[name]
	} else {
		return wire
	}
}

func (wire *Wire) resolve(force bool) int {
	if wire.name[0] == 'x' || wire.name[0] == 'y' {
		return wire.val
	}
	if !force && wire.val > -1 {
		return wire.val
	}
	a := wire.inputA.resolve(force)
	b := wire.inputB.resolve(force)
	wire.val = wire.gate(a, b)

	return wire.val
}

func partOne() {
	finalWires := make([]*Wire, 0)

	for name, wire := range wires {
		if name[0] == 'z' {
			wire.resolve(false)
			finalWires = append(finalWires, wire)
		}
	}

	sort.Slice(finalWires, func(i, j int) bool {
		return finalWires[i].name < finalWires[j].name
	})

	res := uint64(0)
	for i, wire := range finalWires {
		res |= (uint64(wire.val) << i)
	}
	fmt.Println("Part 1:", res)
}

func swapWire(wire1 *Wire, wire2 *Wire) {
	wire1.inputA, wire2.inputA = wire2.inputA, wire1.inputA
	wire1.inputB, wire2.inputB = wire2.inputB, wire1.inputB
	wire1.gate, wire2.gate = wire2.gate, wire1.gate
	wire1.gateName, wire2.gateName = wire2.gateName, wire1.gateName
}

func getWireWithInput(input *Wire) []*Wire {
	res := make([]*Wire, 0)
	for _, wire := range wires {
		if wire.inputA == input || wire.inputB == input {
			res = append(res, wire)
		}
	}
	return res
}

func findCorruptZWires() []*Wire {
	res := make([]*Wire, 0)
	for name, wire := range wires {
		if name[0] != 'z' {
			continue
		}
		if wire.gateName != "XOR" && wire.name != "z45" {
			res = append(res, wire)
			continue
		}
		a, b := wire.inputA, wire.inputB
		if (a.name[0] == 'x' || b.name[0] == 'x') && wire.name != "z00" {
			res = append(res, wire)
		}
	}
	return res
}

func findCorruptFinalXOR() []*Wire {
	res := make([]*Wire, 0)
	for name, wire := range wires {
		if wire.gateName != "XOR" {
			continue
		}
		a, b := wire.inputA, wire.inputB
		if a.name[0] == 'x' || b.name[0] == 'x' {
			continue
		}
		if name[0] != 'z' {
			res = append(res, wire)
		}
	}
	return res
}

func findCorruptXYPar() []*Wire {
	res := make([]*Wire, 0)
	for name, wire := range wires {
		if name[0] == 'x' || name[0] == 'y' || name == "z00" {
			continue
		}
		a, b := wire.inputA, wire.inputB
		if a.name[0] != 'x' && b.name[0] != 'x' {
			continue
		}
		parent := getWireWithInput(wire)
		sort.Slice(parent, func(i, j int) bool {
			return parent[i].gateName < parent[j].gateName
		})
		if wire.gateName == "XOR" {
			if len(parent) != 2 || parent[0].gateName != "AND" || parent[1].gateName != "XOR" {
				res = append(res, wire)
			}
		} else {
			if len(parent) != 1 || parent[0].gateName != "OR" {
				res = append(res, wire)
			}
		}
	}
	return res
}

func dumpWire(wire *Wire, padding int) {
	fmt.Printf("%*s%s = %s %s %s\n", padding, "", wire.name, wire.inputA.name, wire.gateName, wire.inputB.name)
}

func dumpWireInfo(wire *Wire, depth int) {
	dumpWire(wire, (2-depth)*5)
	parent := getWireWithInput(wire)
	if len(parent) > 0 {
		fmt.Printf("%*sAppeared in:\n", (2-depth)*5, "")
		for _, w := range parent {
			if depth == 0 {
				dumpWire(w, (3-depth)*5)
			} else {
				dumpWireInfo(w, depth-1)
			}
		}
	}
}

func flipWire(wire *Wire) {
	if wire.val == -1 {
		return
	}
	wire.val = 1 - wire.val
}

func partTwo() {
	corruptedZ := findCorruptZWires()
	fmt.Println("Corrupted z wires:")
	for _, w := range corruptedZ {
		dumpWireInfo(w, 1)
	}
	fmt.Println()
	// There should be the same number of corrupted final XOR
	corruptedXOR := findCorruptFinalXOR()
	fmt.Println("Corrupted final XOR wires:")
	for _, w := range corruptedXOR {
		dumpWireInfo(w, 1)
		fmt.Println()
	}

	// (Run code to see output, this part should be changed base on output inspection)
	// Found 3 pairs, let's try to swap them back to their correct place
	// z32 should be swapped in gfm, as gfm appeared in a XOR result and an AND result
	// z08 should be swapped with cdj, as its result is used to calculate z09, and currently z08 is calculated using x08 and y08
	// z16 should be swapped with mrb, as they are the only pair left
	swapWire(wires["z32"], wires["gfm"])
	swapWire(wires["z08"], wires["cdj"])
	swapWire(wires["z16"], wires["mrb"])

	// One pair left
	// All z wires are in correct place
	corruptXYPar := findCorruptXYPar()
	fmt.Println("Corrupted x/y parent wires:")
	for _, w := range corruptXYPar {
		dumpWireInfo(w, 1)
		fmt.Println()
	}
	// Found 3, but gsv is there only because the first layer of the adder is special, so there is exactly one pair
	// Swap qjd and dhm then test
	swapWire(wires["qjd"], wires["dhm"])

	res := []string{"z32", "gfm", "z08", "cdj", "z16", "mrb", "qjd", "dhm"}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	fmt.Println("Part 2:", strings.Join(res, ","))

	// Test flipping bit(s) here
	flipWire(wires["x30"])
	flipWire(wires["y25"])

	zWires := make([]*Wire, 0)
	xWires := make([]*Wire, 0)
	yWires := make([]*Wire, 0)

	for name, wire := range wires {
		if name[0] == 'z' {
			wire.resolve(true)
			zWires = append(zWires, wire)
		} else if name[0] == 'x' {
			xWires = append(xWires, wire)
		} else if name[0] == 'y' {
			yWires = append(yWires, wire)
		}
	}

	sort.Slice(zWires, func(i, j int) bool {
		return zWires[i].name < zWires[j].name
	})
	sort.Slice(xWires, func(i, j int) bool {
		return xWires[i].name < xWires[j].name
	})
	sort.Slice(yWires, func(i, j int) bool {
		return yWires[i].name < yWires[j].name
	})

	z := uint64(0)
	for i, wire := range zWires {
		z |= (uint64(wire.val) << i)
	}
	x, y := uint64(0), uint64(0)
	for i := range xWires {
		x |= (uint64(xWires[i].val) << i)
		y |= (uint64(yWires[i].val) << i)
	}
	fmt.Println("Decimal:")
	fmt.Printf("x  =%11d\ny  =%11d\nx+y=%11d\nz  =%11d\n", x, y, x+y, z)
	fmt.Println("Binary:")
	fmt.Printf("x  =%48b\ny  =%48b\nx+y=%48b\nz  =%48b\n", x, y, x+y, z)
}

func main() {
	file, err := os.Open("day-24.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(scanner.Text()) > 0 {
		tokens := strings.Split(scanner.Text(), ": ")
		name := tokens[0]
		val := tokens[1][0] - '0'

		wire := getWire(name)
		wire.val = int(val)
	}

	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		a, b, c := tokens[0], tokens[2], tokens[4]
		gate := gateStr[tokens[1]]

		wireA := getWire(a)
		wireB := getWire(b)
		wireC := getWire(c)

		wireC.inputA = wireA
		wireC.inputB = wireB
		wireC.gate = gate
		wireC.gateName = tokens[1]
	}

	partOne()
	partTwo()
}
