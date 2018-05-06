#!/usr/bin/python3

import statistics
input = open("bus.txt").read().split('\n')[:-1]
# input = ["11 12 14", "2"]

groups = [tuple([int(h)]) for h in input[0].split(" ")]
unhappinessList = [0] * len(groups)
# groups = [[int(h)] for h in input[0].split(" ")]
busstopCount = int(input[1])


def printGrouping(groups):
	for group in groups:
		line = ""
		for h in group:
			line += "%4i " % h
		print(line)


def unhappiness(group):
	unhappinessValue = 0
	if len(group) > 1:
		mean = statistics.mean(group)
		for house in group:
			unhappinessValue += (house - mean) ** 2
	return unhappinessValue


def groupNext(groups, unhappinessList):
	minimum = 10 ** 10
	for i in range(len(groups) - 1):
		u = unhappiness(groups[i] + groups[i + 1]) - unhappinessList[i] - unhappinessList[i + 1]
		if u < minimum:
			minimum = u
			minimumIndex = i
	del(unhappinessList[minimumIndex])
	unhappinessList[minimumIndex] = unhappiness(groups[minimumIndex] + groups[minimumIndex + 1])
	return groups[:minimumIndex] + [groups[minimumIndex] + groups[minimumIndex + 1]] + groups[minimumIndex + 2:], unhappinessList


print(groups)
while len(groups) > busstopCount:
	groups, unhappinessList = groupNext(groups, unhappinessList)
print(groups)
print(sum(unhappinessList))
