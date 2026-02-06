package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run check_cert.go <filename>")
		return
	}

	filename := os.Args[1]
	
	// Print absolute path to help you debug where Go is actually looking
	absPath, _ := filepath.Abs(filename)
	fmt.Printf("🔍 Looking for file at: %s\n", absPath)

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ File Error: %v\n", err)
		return
	}

	// Decode PEM
	block, _ := pem.Decode(data)
	if block == nil {
		fmt.Println("❌ Error: File is not a valid PEM certificate.")
		return
	}

	// Parse Certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("❌ x509 Parser Error: %v\n", err)
		return
	}

	fmt.Printf("✅ Success! Parsed certificate for: %s\n", cert.Subject)
}
