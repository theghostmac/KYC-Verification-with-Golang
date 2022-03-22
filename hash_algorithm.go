package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func showFileToCustomer() {
	fmt.Println("Your document is: " + os.Args[0] + " credential.csv")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		showFileToCustomer()
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	credential := checkArgs()
	// Get bytes from the credential document
	data, err := ioutil.ReadFile(credential)
	if err != nil {
		log.Fatal(err)
	}

	// encrypt the file and display to the customer
	fmt.Printf("The encrypted file is: %x\n\n", sha256.Sum256(data))
}