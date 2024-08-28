package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// Replace 'path/to/your/credentials.json' with the actual path to your downloaded JSON key file.
	credentialsFile := "./tillersheetinteraction-809b737ff165.json"

	// Initialize Google Sheets API
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		log.Fatalf("Unable to initialize Sheets API: %v", err)
	}

	// Now 'srv' is an authenticated client that can interact with Google Sheets API.
	// You can use 'srv' to read, write, or perform other operations on Google Sheets.

	// Example: Get spreadsheet data
	spreadsheetID := "1--kYPkBCIiiCebnum_waqdyzGbRsHClYIRZQw45uGeY"
	readRange := "Transactions!C:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var cumulativeSum float64

	// Print the values from the response
	if len(resp.Values) > 0 {
		fmt.Println("Data from sheet:")
		for i, row := range resp.Values {

			// Remove dollar sign and any commas before parsing
			valueStr := strings.ReplaceAll(row[2].(string), "$", "")
			valueStr = strings.ReplaceAll(valueStr, ",", "")
			// Attempt to parse row[2] as a number and add it to the cumulative sum
			if num, err := strconv.ParseFloat(valueStr, 64); err == nil {
				cumulativeSum += num
			} else {
				log.Printf("Error parsing value in row %d: %v", i+1, err) // Log the error but continue processing
			}
		}
	} else {
		fmt.Println("No data found.")
	}

	log.Println(len(resp.Values))
}
