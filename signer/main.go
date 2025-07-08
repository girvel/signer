package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	signer "github.com/girvel/signer/src"
)

func readPrivateKey(path string) (*rsa.PrivateKey, error) {
    privatePem, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    block, _ := pem.Decode([]byte(privatePem))
    if block == nil {
        return nil, fmt.Errorf("Failed to decode the ./private.pem")
    }

    privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
    if !ok {
        return nil, fmt.Errorf("Key in ./private.pem is not RSA")
    }

    return rsaPrivateKey, nil
}

func main() {
    key, err := readPrivateKey("private.pem")
    if err != nil {
        panic(err.Error())
    }

    api := signer.CreateAPI(key)
    api.Run()
}
