package main

import (
	"fmt"

	"github.com/adnux/go-price-calculator/files"
	"github.com/adnux/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {

		fileName := fmt.Sprintf("result_%.0f.json", taxRate*100)
		fm := files.New("prices.txt", fileName)
		// cm := cmd.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
