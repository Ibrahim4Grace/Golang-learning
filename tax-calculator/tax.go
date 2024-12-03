package main

import (
	"fmt"
	filemanager "tax-calculate/filemanager"
	"tax-calculate/prices"
)

func main() {

	taxRate := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRate))
	errorChans := make([]chan error, len(taxRate))

	// //4loops
	// for _, taxRate := range taxRate {
	// 	fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
	// 	priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
	// 	// err := priceJob.Process()
	// 	// if err != nil {
	// 	// 	fmt.Println("Could not process job")
	// 	// 	fmt.Println(err)
	// 	// }

	// }
	//To speed up proccessing using gorouting
	for index, taxRate := range taxRate {
		//To speed up proccessing
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		// proccess will let us know through the channel when its done
		go priceJob.Process(doneChans[index], errorChans[index])

	}

	//handling error and sucess through gorouting
	for index := range taxRate {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}

		case <-doneChans[index]:
			fmt.Println("Done")

		}
	}

}

//To include delays
// time.Sleep(1 * time.Second)
