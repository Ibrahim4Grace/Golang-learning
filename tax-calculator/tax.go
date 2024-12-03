package main

import (
	"fmt"
	filemanager "tax-calculate/filemanager"
	"tax-calculate/prices"
)

func main() {

	taxRate := []float64{0, 0.07, 0.1, 0.15}

	//4loops
	for _, taxRate := range taxRate {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

	}
}
