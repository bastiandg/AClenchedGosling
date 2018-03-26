#!/usr/bin/python3

import sys

# path = [[  4, "*",  11],
	# ["+",   4, "-"],
	# [ 22, "-",   9]]

path = [["*",   8, "-",   1],
	[  4, "*",  11, "*"],
	["+",   4, "-",  18],
	[ 22, "-",   9, "*"]]


def moveNorth(x, y):
	if y == 0 or (x == len(path) - 1 and y == 0):
		return x, y, False
	else:
		return x, y - 1, True


def moveSouth(x, y):
	if y == len(path) - 1 or (y == len(path) - 2 and x == 0) or (x == len(path) - 1 and y == 0):
		return x, y, False
	else:
		return x, y + 1, True


def moveEast(x, y):
	if x == len(path[0]) - 1 or (x == len(path) - 1 and y == 0):
		return x, y, False
	else:
		return x + 1, y, True


def moveWest(x, y):
	if x == 0 or (x == 1 and y == len(path) - 1) or (x == len(path) - 1 and y == 0):
		return x, y, False
	else:
		return x - 1, y, True


def findPath(moveList, x, y, value):
	if x == (len(path[0]) - 1) and y == 0 and value == targetValue:
		print(moveList.replace("E", "O"))
		return moveList
	elif len(moveList) >= maxSteps:
		return "error"
	for step1 in ["N", "S", "E", "W"]:
		x2, y2, success = move[step1](x, y)
		if not success:
			continue
		operator = path[y2][x2]
		for step2 in ["N", "S", "E", "W"]:
			x3, y3, success = move[step2](x2, y2)
			if not success:
				continue
			operand = path[y3][x3]
			if operator == "+":
				tmpValue = value + operand
			elif operator == "-":
				tmpValue = value - operand
			elif operator == "*":
				tmpValue = value * operand
			elif operator == "/":
				tmpValue = value / operand
			findPath(moveList + step1 + step2, x3, y3, tmpValue)


y = len(path) - 1
x = 0
targetValue = 30
maxSteps = 12
move = {"N": moveNorth,
	"S": moveSouth,
	"E": moveEast,
	"W": moveWest}

findPath("", x, y, 22)
