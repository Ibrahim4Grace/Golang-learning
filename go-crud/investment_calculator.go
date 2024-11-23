package main // for go to know where app started

import (
	"fmt"
	"math"
)
const inflationRate = 2.5

//we use float64 for decimal numbers
//we can use : instead of var to set veriables
//math.Pow is used to calculate
 
func main() {

var invesmentAmount float64
var years float64 
 expectedReturnRate := 5.5

//to show something for the user
fmt.Print("Enter investment amount: ")
// scan to get data from user
fmt.Scan(&invesmentAmount)

fmt.Print("Expected return rate: ")
fmt.Scan(&expectedReturnRate)

fmt.Print("Enter investment years: ")
fmt.Scan(&years)

// we can also use out function in here since it return 2 veriables
futureValue,futureRealValue := calculateFutureValues(invesmentAmount, expectedReturnRate, years)

//  futureValue := invesmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
//  futureRealValue := futureValue / math.Pow(1 + inflationRate/100 , years)



// output answer
// fmt.Println("Future value: ", futureValue)
// fmt.Println("Future value adjusted for inflation", futureRealValue)

//to format output information dot 2f stands for decimal point we want 
// fmt.Printf("Future value: %.2f\nFuture value adjusted for inflation: %.2f", futureValue,futureRealValue)

//we can also format like this 
formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
formattedRFV := fmt.Sprintf("Future value adjusted for inflation: %.2f\n",futureRealValue)

fmt.Print(formattedFV, formattedRFV)

}

//how to use function in Golang 
func calculateFutureValues(invesmentAmount, expectedReturnRate, years float64)(fv float64, rfv float64){
	fv = invesmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, rfv
}