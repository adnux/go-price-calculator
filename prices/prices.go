package prices

import (
	"fmt"

	"github.com/adnux/go-price-calculator/conversion"
	"github.com/adnux/go-price-calculator/ioManager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         ioManager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(iom ioManager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool) {
	err := job.LoadData()

	if err != nil {
		// return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)

	fmt.Println(result)

	doneChannel <- true

	// return nil
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices

	return nil
}
