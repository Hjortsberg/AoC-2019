package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

func newCoordinate(x int, y int) coordinate {
	c := coordinate{x: x}
	c.y = y
	return c
}

func main() {
	var cableOneVisitedCoordinates []coordinate
	var cableTwoVisitedCoordinates []coordinate
	cableOneVisitedCoordinates = append(cableOneVisitedCoordinates, newCoordinate(0, 0))
	cableTwoVisitedCoordinates = append(cableTwoVisitedCoordinates, newCoordinate(0, 0))

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

	nextCable := false
	for scanner.Scan() {
		cableDirection := scanner.Text()
		if cableDirection == "\n" {
			nextCable = true
			fmt.Println("switched cable")
			continue
		}
		if nextCable == false {
			cableOneVisitedCoordinates = traceCable(cableDirection, cableOneVisitedCoordinates)
		} else {
			cableTwoVisitedCoordinates = traceCable(cableDirection, cableTwoVisitedCoordinates)
		}
	}

	intersections := findIntersections(cableOneVisitedCoordinates, cableTwoVisitedCoordinates)
	fmt.Println(manhattanDist(intersections))
}

func manhattanDist(intersections []coordinate) int {
	smallestManhattanDist := 100000
	for _, c := range intersections[1:] {
		pVecDst := 0 - c.x
		if pVecDst < 0 {
			pVecDst = -pVecDst
		}
		qVecDst := 0 - c.y
		if qVecDst < 0 {
			qVecDst = -qVecDst
		}
		dst := pVecDst + qVecDst
		if dst < smallestManhattanDist {
			smallestManhattanDist = dst
		}
	}
	return smallestManhattanDist
}

func findIntersections(cableOne []coordinate, cableTwo []coordinate) []coordinate {
	var intersectingCordinates []coordinate
	shortestPath := 1000000
	for s1, cordCO := range cableOne {
		for s2, cordCT := range cableTwo {
			if cordCO == cordCT {
				intersectingCordinates = append(intersectingCordinates, cordCO)
				if (s1+s2) < shortestPath && (s1+s2) != 0 {
					shortestPath = s1 + s2
				}
			}
		}
	}
	fmt.Println("shortestPath")
	fmt.Println(shortestPath)
	return intersectingCordinates
}

func traceCable(direction string, currCable []coordinate) []coordinate {
	lastCoordinate := currCable[0]
	if len(currCable) >= 1 {
		lastCoordinate = currCable[len(currCable)-1]
	}
	numSteps, err := strconv.Atoi(direction[1:])
	if err != nil {
		fmt.Println(err)
		fmt.Println("error converting cable steps")
	}
	switch direction[0] {
	case 10:
	case 82:
		for i := 1; i <= numSteps; i++ {
			currCable = append(currCable, newCoordinate(lastCoordinate.x+i, lastCoordinate.y))
		}
	case 76:
		for i := 1; i <= numSteps; i++ {
			currCable = append(currCable, newCoordinate(lastCoordinate.x-i, lastCoordinate.y))
		}
	case 85:
		for i := 1; i <= numSteps; i++ {
			currCable = append(currCable, newCoordinate(lastCoordinate.x, lastCoordinate.y+i))
		}
	case 68:
		for i := 1; i <= numSteps; i++ {
			currCable = append(currCable, newCoordinate(lastCoordinate.x, lastCoordinate.y-i))
		}
	default:
		fmt.Println("an error occured in the switchcase")
	}
	return currCable
}
