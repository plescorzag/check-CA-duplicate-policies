package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run check_bundle.go <bundle-file.crt>")
		return
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ File Error: %v\n", err)
		return
	}

	count := 0
	errorsFound := 0
	rest := data

	fmt.Printf("Testing bundle: %s\n", filename)
	fmt.Println("--------------------------------------------------")

	for {
		var block *pem.Block
		block, rest = pem.Decode(rest) // Decode one block and keep the 'rest' for the next loop
		if block == nil {
			break
		}

		if block.Type != "CERTIFICATE" {
			continue
		}

		count++
		fmt.Printf("[%d] Subject: ", count)

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			fmt.Printf("INVALID\n    ❌ Error: %v\n", err)
			errorsFound++
		} else {
			fmt.Printf("%s\n    ✅ OK\n", cert.Subject.CommonName)
		}

		if len(rest) == 0 {
			break
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Finished. Processed %d certificates with %d errors.\n", count, errorsFound)
}
