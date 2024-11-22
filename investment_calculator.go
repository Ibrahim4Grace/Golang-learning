package main // for go to know where app started

import (
	"fmt"
	"math"
)

func main() {
var invesmentAmount float64 = 1000
var expectedReturnRate = 5.5 
var years float64 = 10

var futureValue = invesmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

// output answer
fmt.Println(futureValue)
}

