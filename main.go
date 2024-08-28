package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {


	// Now 'srv' is an authenticated client that can interact with Google Sheets API.
	// You can use 'srv' to read, write, or perform other operations on Google Sheets.

	// Example: Get spreadsheet data
	readRange := "Transactions!C:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var cumulativeSum float64

	// Print the values from the response
	if len(resp.Values) > 0 {
		fmt.Println("Data from sheet:")
		for _, row := range resp.Values {

			// Remove dollar sign and any commas before parsing
			valueStr := strings.ReplaceAll(row[2].(string), "$", "")
			valueStr = strings.ReplaceAll(valueStr, ",", "")
			// Attempt to parse row[2] as a number and add it to the cumulative sum
			if num, err := strconv.ParseFloat(valueStr, 64); err == nil {
				cumulativeSum += num
			}
		}
	} else {
		fmt.Println("No data found.")
	}

	log.Println('$', cumulativeSum)
}
