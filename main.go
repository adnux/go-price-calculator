package main

import (
	"fmt"
	"os"

	"github.com/adnux/go-price-calculator/files"
	"github.com/adnux/go-price-calculator/prices"
)

func DeletePreviousGeneratedFiles() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if len(file.Name()) > 7 && file.Name()[:7] == "result_" {
			err := os.Remove(file.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	// errorChannels := make([]chan error, len(taxRates))

	err := DeletePreviousGeneratedFiles()
	if err != nil {
		fmt.Println("Could not delete previous generated files")
		fmt.Println(err)
	}

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)

		fileName := fmt.Sprintf("result_%.0f.json", taxRate*100)
		fm := files.New("prices.txt", fileName)
		// cm := cmd.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.Process(doneChannels[index])
		go priceJob.Process(doneChannels[index])

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

	for _, doneChannel := range doneChannels {
		<-doneChannel
	}
}
