package prices

import (
	"fmt"
	"tax-calculate/conversion"
	"tax-calculate/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager        filemanager.FileManager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	InputPrice       []float64               `json:"input_price"`
	TaxIncludePrices map[string]string       `json:"tax_include_prices"`
}

// To read prices from file .txt
func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	// importing from pkg
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrice = prices
	return nil
}

// adding reciever method before function name
func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	//sending error through channels
	if err != nil {
		// return err
		errorChan <- err
		return
	}
	result := make(map[string]string)
	for _, price := range job.InputPrice {

		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	// fmt.Println(result)
	//OR WRITE TOT FILE
	job.TaxIncludePrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true

}

// using construction function
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:  fm,
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}

}
