package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GSheetClient struct {
	srv *sheets.Service
}

func NewGSheetClient() (*GSheetClient, error) {

	err := godotenv.Load("./.env")
	if err != nil {
		return nil, err
	}

	// credentialsFile := os.Getenv("CREDENTIALS_FILE")
	// if credentialsFile == "" {
	// 	return nil,
	// 	log.Fatal("CREDENTIALS_FILE environment variable is not set")
	// }

	spreadsheetID := os.Getenv("SPREADSHEET_ID")
	if spreadsheetID == "" {
		log.Fatal("SPREADSHEET_ID environment variable is not set")
	}

	// Initialize Google Sheets API
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		log.Fatalf("Unable to initialize Sheets API: %v", err)
	}

	return &GSheetClient{srv: srv}, nil
}
