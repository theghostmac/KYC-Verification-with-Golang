package main

import (
	"log"
	"os"
)

// openCustomerID opens the credentials
func openCustomerID() {
	customerID, err := os.Open("credential.csv")
	if err != nil {
		log.Fatal(err)
	}
}g

func RetrieveCustomerID() {
	// collect the credential.csv file and run QueryCustomerID()
}

func QueryCustomerID() {
	// return decrypted value of the credential.csv file and compare
}

func main() {
	defer file.Close()
}
