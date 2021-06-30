package main

import (
	_ "bytes"
)

func maxOnesAfterRemoveItem(bitArray []byte) uint {
	var sumLeft uint = 0
	var sumRight uint = 0
	var maxOnes uint = 0
	var zeroExpected bool = false

	for i := 0; i < len(bitArray); i++ {
		if !zeroExpected && i == len(bitArray)-1 && bitArray[i] == 1 {
			continue
		}
		if bitArray[i] == 1 {
			sumRight++
		} else {
			zeroExpected = true
			if sumLeft+sumRight > maxOnes {
				maxOnes = sumLeft + sumRight
			}
			sumLeft = sumRight
			sumRight = 0
		}
	}

	if sumLeft+sumRight > maxOnes {
		maxOnes = sumLeft + sumRight
	}

	return maxOnes
}

func main() {
}
