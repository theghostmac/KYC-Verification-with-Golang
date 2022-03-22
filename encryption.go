package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// read the customer's credential document saved as customer.pdf
	customerID, err := ioutil.ReadFile("customer.pdf")
	// encrypt the file
	encryptFile, err := encrypt(customerID)
	if err != nil {
		log.Fatal(err)
	}
	// display the encrypted output to the customer
	fmt.Printf("Sha256 Hashing: %x\n\n", sha256.Sum256(customerID))
}
