package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"./business_hours"
)

func main() {
	csvFile, err := os.Open("./question3.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	var row []string
	var binResult string
	for {
		row, err = reader.Read()
		if err == io.EOF {
			break
		}

		ret := business_hours.IsValidBusinessHours(row)
		fmt.Printf("%s -> %v\n", row, ret)

		if ret == true {
			binResult += "0"
		} else {
			binResult += "1"
		}
	}

	fmt.Println("---")
	fmt.Println("Binary Result: ", binResult)

	result, _ := strconv.ParseInt(binResult, 2, 64)
	fmt.Println("Final Answer: ", result)
}
