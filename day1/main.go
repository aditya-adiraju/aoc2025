package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var FILE_PATH string

func mod(a, b int) int {
	return (a % b + b) % b
}

type Rotation struct {
	rot byte 
	clicks int	
}

func getRotations() []Rotation {
	rotation_list := []Rotation{};

	fi, err := os.ReadFile(FILE_PATH)
	input_string := string(fi[:])
	if err != nil {
		panic(err)
	}

	rots := strings.Split(input_string, "\n");

	for _, v := range rots {
		mag, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)	
		}

		if !(v[0] == 'R' || v[0] == 'L') { 
			panic("first character is not L or R")
		}

		rotation_list = append(rotation_list, Rotation{v[0], mag})
	}

	return rotation_list
}

func part1() {
	rotation_list := getRotations();
	arrow := 50
	password := 0

	for _, rot := range rotation_list {
		if rot.rot == 'R' {
			arrow += rot.clicks
		} else {
			arrow -= rot.clicks
		}

		if mod(arrow, 100) == 0 {
			password++
		}
	}

	fmt.Println("part 1 password:", password)
}


func part2() {

	rotation_list := getRotations();
	arrow := 50
	password := 0
	for _, rot := range rotation_list {
		// adjust arrow
		if rot.rot == 'R' {
			arrow += rot.clicks
		} else {
			arrow -= rot.clicks
		}

		// increment password based on arrow value
		if arrow == 0 {
			password++
		} else if arrow < 0 {
			password += (-1 * (arrow / 100))
			if -1 * rot.clicks != arrow {
				password++
			}
		} else {
			password +=  (arrow / 100)
		}

		// reset arrow to true value
		arrow = mod(arrow, 100);

		// fmt.Println(rot, arrow, password)
	}
	fmt.Println("part 2 password:", password)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "TEST" {
		FILE_PATH = "./test.txt"
	} else {
		FILE_PATH = "./input.txt"
	}

	part1();
	part2();
}