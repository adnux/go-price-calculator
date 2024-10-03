package cmd

import "fmt"

type CmdManager struct {
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println()
	var prices []string

	for {
		fmt.Print("Enter price:")
		var price string
		fmt.Scanln(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CmdManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CmdManager {
	return CmdManager{}
}
