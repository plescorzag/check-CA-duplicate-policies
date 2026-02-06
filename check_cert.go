package main

import (
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "os"
)

func main() {
    // Read the CA certificate file
    data, err := os.ReadFile("ca.crt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    // Decode the PEM block
    block, _ := pem.Decode(data)
    if block == nil {
        fmt.Println("Failed to decode PEM block. Is it a valid certificate file?")
        return
    }

    // Attempt to parse the certificate
    _, err = x509.ParseCertificate(block.Bytes)
    if err != nil {
        fmt.Printf("❌ Go Parser Error: %v\n", err)
    } else {
        fmt.Println("✅ Certificate parsed successfully!")
    }
}
