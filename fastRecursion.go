package main

import "fmt"
import "strings"

type Move struct {
	x       int
	y       int
	success bool
}

const targetValue = 30

var resultFound = false

/*
-1 +
-2 -
-3 *
-4 /
*/

var path = [4][4]int{
	{-3, 8, -2, 1},
	{4, -3, 11, -3},
	{-1, 4, -2, 18},
	{22, -2, 9, -3}}

var steps = [4]string{"N", "S", "E", "W"}

func moveNorth(x int, y int) Move {
	if y == 0 || (x == len(path)-1 && y == 0) {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x, y: y - 1, success: true}
	}
}

func moveSouth(x int, y int) Move {
	if y == len(path)-1 || (y == len(path)-2 && x == 0) || (x == len(path)-1 && y == 0) {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x, y: y + 1, success: true}
	}
}

func moveEast(x int, y int) Move {
	if x == len(path[0])-1 || (x == len(path)-1 && y == 0) {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x + 1, y: y, success: true}
	}
}

func moveWest(x int, y int) Move {
	if x == 0 || (x == 1 && y == len(path)-1) || (x == len(path)-1 && y == 0) {
		return Move{x: x, y: y, success: false}
	} else {
		return Move{x: x - 1, y: y, success: true}
	}
}

func findPath(moveList string, x int, y int, value int, maxSteps int) string {
	if x == (len(path[0])-1) && y == 0 && value == targetValue {
		resultFound = true
		return moveList
	} else if len(moveList) >= maxSteps {
		return "error"
	}
	for i := 0; i < 4; i++ { // step1 in ["N", "S", "E", "W"]:
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
		for j := 0; j < 4; j++ { // step1 in ["N", "S", "E", "W"]:
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
			if resultFound{
				return result
			}
		}
	}
	return "error"
}

func main() {
	x := 0
	y := len(path) - 1
	for maxSteps := 8; !resultFound; maxSteps++ {
		result := findPath("", x, y, path[y][x], maxSteps)
		if result != "error"{
			fmt.Println(strings.Replace(result, "E", "O", -1))
		}
	}
}
