package main

import "fmt"
import "strings"
import "io/ioutil"
import "strconv"

type Move struct {
	x       int
	y       int
	success bool
}

var targetValue int
var path [][]int
var steps = [4]string{"N", "S", "E", "W"}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func readPath() [][]int{
	fileBytes, err := ioutil.ReadFile("exampleinput_5_hyper.txt")
	check(err)
	file := string(fileBytes)
	file = strings.Replace(file, "-", "-2", -1)
	file = strings.Replace(file, "+", "-1", -1)
	file = strings.Replace(file, "*", "-3", -1)
	file = strings.Replace(file, "/", "-4", -1)

	lines := strings.Split(file, "\n")
	targetValue, err = strconv.Atoi(lines[0])
	check(err)
	length, err := strconv.Atoi(lines[1])
	check(err)
	square := lines[2:]

	path2 := make([][]int, length)
	var squareLine []string
	for i := range path2 {
		squareLine = strings.Split(square[i], ";")
		path2[i] = make([]int, length)
		for j := range path2[i] {
			path2[i][j], err = strconv.Atoi(squareLine[j])
			check(err)
		}
	}
	return path2
}


func moveNorth(x int, y int) Move {
	if y == 0 {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x, y: y - 1, success: true}
	}
}

func moveSouth(x int, y int) Move {
	if y == len(path) - 1 {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x, y: y + 1, success: true}
	}
}

func moveEast(x int, y int) Move {
	if x == len(path[0]) - 1 {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x + 1, y: y, success: true}
	}
}

func moveWest(x int, y int) Move {
	if x == 0 || (y == 0 && x == 1) {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x - 1, y: y, success: true}
	}
}

func findPath(moveList string, x int, y int, value int, maxSteps int) bool {
	if x == (len(path[0])-1) && y == (len(path[0])-1) && value == targetValue {
		fmt.Println(strings.Replace(moveList, "E", "O", -1))
		return true
	} else if len(moveList) >= maxSteps {
		return false
	}
	for i := 0; i < 4; i++ {
		step1 := steps[i]
		var move1 Move
		switch step1 {
		case "N":
			move1 = moveNorth(x, y)
		case "S":
			move1 = moveSouth(x, y)
		case "E":
			move1 = moveEast(x, y)
		case "W":
			move1 = moveWest(x, y)
		}
		if !move1.success {
			continue
		}
		operator := path[move1.y][move1.x]
		for j := 0; j < 4; j++ {
			step2 := steps[j]
			var move2 Move
			switch step2 {
			case "N":
				move2 = moveNorth(move1.x, move1.y)
			case "S":
				move2 = moveSouth(move1.x, move1.y)
			case "E":
				move2 = moveEast(move1.x, move1.y)
			case "W":
				move2 = moveWest(move1.x, move1.y)
			}
			if !move2.success {
				continue
			}
			operand := path[move2.y][move2.x]
			var tmpValue int
			switch operator {
			case -1:
				tmpValue = value + operand
			case -2:
				tmpValue = value - operand
			case -3:
				tmpValue = value * operand
			case -4:
				tmpValue = value / operand
			}
			result := findPath(moveList+step1+step2, move2.x, move2.y, tmpValue, maxSteps)
			if result {
				return result
			}
		}
	}
	return false
}

func main() {
	x := 0
	y := 0
	path = readPath()
	var resultFound = false
	for maxSteps := 8; !resultFound; maxSteps++ {
		resultFound = findPath("", x, y, path[y][x], maxSteps)
	}
}
