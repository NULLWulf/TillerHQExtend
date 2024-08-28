package main

type Transaction struct {
	Date            string
	Description     string
	Category        string
	Amount          string
	Account         string
	AccountNumber   string
	Institution     string
	Month           string
	Week            string
	TransactionID   string
	AccountID       string
	CheckNumber     string
	FullDescription string
	DateAdded       string
}

func MapTransaction(row []interface{}) Transaction {
	return Transaction{
		Date:            row[0].(string),
		Description:     row[1].(string),
		Category:        row[2].(string),
		Amount:          row[3].(string),
		Account:         row[4].(string),
		AccountNumber:   row[5].(string),
		Institution:     row[6].(string),
		Month:           row[7].(string),
		Week:            row[8].(string),
		TransactionID:   row[9].(string),
		AccountID:       row[10].(string),
		CheckNumber:     row[11].(string),
		FullDescription: row[12].(string),
		DateAdded:       row[13].(string),
	}
}
