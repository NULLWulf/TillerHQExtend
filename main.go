package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	// Print the values from the response
	if len(resp.Values) > 0 {
		fmt.Println("Data from sheet:")
		for i, row := range resp.Values {
			fmt.Printf("%s, %s\n", row[0], row[2])
			if i == 2 {
				os.Exit(5)
			}
		}
	} else {
		fmt.Println("No data found.")
	}

	log.Println(len(resp.Values))
}
